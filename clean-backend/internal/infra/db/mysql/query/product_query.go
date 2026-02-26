package query

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/tokushun109/tku/clean-backend/internal/infra/db/mysql/mysqlutil"
	usecaseProductQuery "github.com/tokushun109/tku/clean-backend/internal/usecase/product/query"
)

type ProductQueryReader struct {
	db *sqlx.DB
}

type productBaseRow struct {
	ID           uint           `db:"id"`
	UUID         string         `db:"uuid"`
	Name         string         `db:"name"`
	Description  sql.NullString `db:"description"`
	Price        int            `db:"price"`
	IsActive     bool           `db:"is_active"`
	IsRecommend  bool           `db:"is_recommend"`
	CategoryUUID sql.NullString `db:"category_uuid"`
	CategoryName sql.NullString `db:"category_name"`
	TargetUUID   sql.NullString `db:"target_uuid"`
	TargetName   sql.NullString `db:"target_name"`
}

type productTagRow struct {
	ProductID uint   `db:"product_id"`
	TagUUID   string `db:"tag_uuid"`
	TagName   string `db:"tag_name"`
}

type productImageRow struct {
	ProductID uint   `db:"product_id"`
	UUID      string `db:"uuid"`
	Name      string `db:"name"`
	MimeType  string `db:"mime_type"`
	Path      string `db:"path"`
	Order     int    `db:"order"`
}

type productSiteDetailRow struct {
	ProductID     uint   `db:"product_id"`
	UUID          string `db:"uuid"`
	DetailURL     string `db:"detail_url"`
	SalesSiteUUID string `db:"sales_site_uuid"`
	SalesSiteName string `db:"sales_site_name"`
}

type categoryRow struct {
	UUID string `db:"uuid"`
	Name string `db:"name"`
}

func NewProductQueryReader(db *sqlx.DB) *ProductQueryReader {
	return &ProductQueryReader{db: db}
}

func (r *ProductQueryReader) ListProducts(ctx context.Context, q usecaseProductQuery.ListProductsQuery) ([]*usecaseProductQuery.Product, error) {
	query := `
		SELECT
			p.id,
			p.uuid,
			p.name,
			p.description,
			p.price,
			p.is_active,
			p.is_recommend,
			c.uuid AS category_uuid,
			c.name AS category_name,
			t.uuid AS target_uuid,
			t.name AS target_name
		FROM product p
		LEFT JOIN category c ON c.id = p.category_id AND c.deleted_at IS NULL
		LEFT JOIN target t ON t.id = p.target_id AND t.deleted_at IS NULL
		WHERE p.deleted_at IS NULL
	`
	args := make([]any, 0, 2)
	if q.Mode == "active" {
		query += ` AND p.is_active = 1`
	}
	if q.Category != "all" {
		query += ` AND c.uuid = ?`
		args = append(args, q.Category)
	}
	if q.Target != "all" {
		query += ` AND t.uuid = ?`
		args = append(args, q.Target)
	}
	query += ` ORDER BY p.created_at DESC`

	rows := []productBaseRow{}
	if err := r.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, err
	}

	products := buildProductsFromRows(rows)
	if len(products) == 0 {
		return products, nil
	}

	if err := r.fillChildren(ctx, products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductQueryReader) ListCategoryProducts(ctx context.Context, q usecaseProductQuery.ListCategoryProductsQuery) ([]*usecaseProductQuery.CategoryProducts, error) {
	products, err := r.ListProducts(ctx, usecaseProductQuery.ListProductsQuery{
		Mode:     q.Mode,
		Category: q.Category,
		Target:   q.Target,
	})
	if err != nil {
		return nil, err
	}

	if q.Category != "all" {
		foundCategory, found, err := r.findCategoryByUUID(ctx, q.Category)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, usecaseProductQuery.ErrCategoryNotFound
		}

		return []*usecaseProductQuery.CategoryProducts{
			{
				Category: foundCategory,
				Products: products,
			},
		}, nil
	}

	categories, err := r.listUsedCategories(ctx)
	if err != nil {
		return nil, err
	}
	if len(categories) == 0 {
		return []*usecaseProductQuery.CategoryProducts{}, nil
	}

	productsByCategoryUUID := make(map[string][]*usecaseProductQuery.Product, len(categories))
	for _, product := range products {
		productsByCategoryUUID[product.Category.UUID] = append(productsByCategoryUUID[product.Category.UUID], product)
	}

	result := make([]*usecaseProductQuery.CategoryProducts, 0, len(categories))
	for _, category := range categories {
		result = append(result, &usecaseProductQuery.CategoryProducts{
			Category: category,
			Products: productsByCategoryUUID[category.UUID],
		})
	}

	return result, nil
}

func (r *ProductQueryReader) ListCarouselItems(ctx context.Context, q usecaseProductQuery.ListCarouselQuery) ([]*usecaseProductQuery.CarouselItem, error) {
	if q.Limit <= 0 {
		return []*usecaseProductQuery.CarouselItem{}, nil
	}

	recommendedRows, err := r.listCarouselBaseRows(ctx, q.Limit, true, nil)
	if err != nil {
		return nil, err
	}

	baseRows := recommendedRows
	if len(baseRows) < q.Limit {
		remain := q.Limit - len(baseRows)

		excludeProductIDs := make([]uint, 0, len(baseRows))
		for _, row := range baseRows {
			excludeProductIDs = append(excludeProductIDs, row.ID)
		}

		newRows, err := r.listCarouselBaseRows(ctx, remain, false, excludeProductIDs)
		if err != nil {
			return nil, err
		}
		baseRows = append(baseRows, newRows...)
	}

	products := buildProductsFromRows(baseRows)
	if len(products) == 0 {
		return []*usecaseProductQuery.CarouselItem{}, nil
	}

	if err := r.fillChildren(ctx, products); err != nil {
		return nil, err
	}

	items := make([]*usecaseProductQuery.CarouselItem, 0, len(products))
	for _, product := range products {
		items = append(items, &usecaseProductQuery.CarouselItem{Product: product})
	}

	return items, nil
}

func (r *ProductQueryReader) GetProductByUUID(ctx context.Context, productUUID string) (*usecaseProductQuery.Product, error) {
	row := productBaseRow{}
	err := r.db.GetContext(
		ctx,
		&row,
		`
		SELECT
			p.id,
			p.uuid,
			p.name,
			p.description,
			p.price,
			p.is_active,
			p.is_recommend,
			c.uuid AS category_uuid,
			c.name AS category_name,
			t.uuid AS target_uuid,
			t.name AS target_name
		FROM product p
		LEFT JOIN category c ON c.id = p.category_id AND c.deleted_at IS NULL
		LEFT JOIN target t ON t.id = p.target_id AND t.deleted_at IS NULL
		WHERE p.uuid = ? AND p.deleted_at IS NULL
		`,
		productUUID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	products := buildProductsFromRows([]productBaseRow{row})
	if err := r.fillChildren(ctx, products); err != nil {
		return nil, err
	}

	return products[0], nil
}

func buildProductsFromRows(rows []productBaseRow) []*usecaseProductQuery.Product {
	products := make([]*usecaseProductQuery.Product, 0, len(rows))
	for _, row := range rows {
		product := &usecaseProductQuery.Product{
			ID:          row.ID,
			UUID:        row.UUID,
			Name:        row.Name,
			Description: mysqlutil.NullStringOrEmpty(row.Description),
			Price:       row.Price,
			IsActive:    row.IsActive,
			IsRecommend: row.IsRecommend,
			Category: usecaseProductQuery.Classification{
				UUID: mysqlutil.NullStringOrEmpty(row.CategoryUUID),
				Name: mysqlutil.NullStringOrEmpty(row.CategoryName),
			},
			Target: usecaseProductQuery.Classification{
				UUID: mysqlutil.NullStringOrEmpty(row.TargetUUID),
				Name: mysqlutil.NullStringOrEmpty(row.TargetName),
			},
			Tags:          []usecaseProductQuery.Classification{},
			ProductImages: []usecaseProductQuery.ProductImage{},
			SiteDetails:   []usecaseProductQuery.SiteDetail{},
		}
		products = append(products, product)
	}
	return products
}

func (r *ProductQueryReader) fillChildren(ctx context.Context, products []*usecaseProductQuery.Product) error {
	if len(products) == 0 {
		return nil
	}

	productIDs := make([]uint, 0, len(products))
	productIndex := make(map[uint]int, len(products))
	for i, product := range products {
		productIDs = append(productIDs, product.ID)
		productIndex[product.ID] = i
	}

	tagsByProductID, err := r.loadTags(ctx, productIDs)
	if err != nil {
		return err
	}
	imagesByProductID, err := r.loadImages(ctx, productIDs)
	if err != nil {
		return err
	}
	siteDetailsByProductID, err := r.loadSiteDetails(ctx, productIDs)
	if err != nil {
		return err
	}

	for productID, tags := range tagsByProductID {
		products[productIndex[productID]].Tags = tags
	}
	for productID, images := range imagesByProductID {
		products[productIndex[productID]].ProductImages = images
	}
	for productID, siteDetails := range siteDetailsByProductID {
		products[productIndex[productID]].SiteDetails = siteDetails
	}

	return nil
}

func (r *ProductQueryReader) loadTags(ctx context.Context, productIDs []uint) (map[uint][]usecaseProductQuery.Classification, error) {
	query, args, err := sqlx.In(
		`SELECT ptt.product_id, t.uuid AS tag_uuid, t.name AS tag_name
		 FROM product_to_tag ptt
		 INNER JOIN tag t ON t.id = ptt.tag_id AND t.deleted_at IS NULL
		 WHERE ptt.deleted_at IS NULL AND ptt.product_id IN (?)
		 ORDER BY ptt.id ASC`,
		productIDs,
	)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)

	rows := []productTagRow{}
	if err := r.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, err
	}

	result := make(map[uint][]usecaseProductQuery.Classification, len(productIDs))
	for _, row := range rows {
		result[row.ProductID] = append(result[row.ProductID], usecaseProductQuery.Classification{
			UUID: row.TagUUID,
			Name: row.TagName,
		})
	}
	return result, nil
}

func (r *ProductQueryReader) loadImages(ctx context.Context, productIDs []uint) (map[uint][]usecaseProductQuery.ProductImage, error) {
	query, args, err := sqlx.In(
		"SELECT pi.product_id, pi.uuid, pi.name, pi.mime_type, pi.path, pi.`order` AS `order`\n"+
			"FROM product_image pi\n"+
			"WHERE pi.deleted_at IS NULL AND pi.product_id IN (?)\n"+
			"ORDER BY pi.`order` DESC, pi.id ASC",
		productIDs,
	)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)

	rows := []productImageRow{}
	if err := r.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, err
	}

	result := make(map[uint][]usecaseProductQuery.ProductImage, len(productIDs))
	for _, row := range rows {
		result[row.ProductID] = append(result[row.ProductID], usecaseProductQuery.ProductImage{
			UUID:     row.UUID,
			Name:     row.Name,
			MimeType: row.MimeType,
			Path:     row.Path,
			Order:    row.Order,
			APIPath:  "",
		})
	}
	return result, nil
}

func (r *ProductQueryReader) loadSiteDetails(ctx context.Context, productIDs []uint) (map[uint][]usecaseProductQuery.SiteDetail, error) {
	query, args, err := sqlx.In(
		`SELECT
			sd.product_id,
			sd.uuid,
			sd.detail_url,
			ss.uuid AS sales_site_uuid,
			ss.name AS sales_site_name
		 FROM site_detail sd
		 INNER JOIN sales_site ss ON ss.id = sd.sales_site_id AND ss.deleted_at IS NULL
		 WHERE sd.deleted_at IS NULL AND sd.product_id IN (?)
		 ORDER BY sd.detail_url ASC, sd.id ASC`,
		productIDs,
	)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)

	rows := []productSiteDetailRow{}
	if err := r.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, err
	}

	result := make(map[uint][]usecaseProductQuery.SiteDetail, len(productIDs))
	for _, row := range rows {
		result[row.ProductID] = append(result[row.ProductID], usecaseProductQuery.SiteDetail{
			UUID:      row.UUID,
			DetailURL: row.DetailURL,
			SalesSite: usecaseProductQuery.SalesSite{
				UUID: row.SalesSiteUUID,
				Name: row.SalesSiteName,
			},
		})
	}
	return result, nil
}

func (r *ProductQueryReader) listUsedCategories(ctx context.Context) ([]usecaseProductQuery.Classification, error) {
	rows := []categoryRow{}
	query := `
		SELECT c.uuid, c.name
		FROM category c
		INNER JOIN product p ON p.category_id = c.id AND p.deleted_at IS NULL
		WHERE c.deleted_at IS NULL
		GROUP BY c.id, c.uuid, c.name
		ORDER BY c.id ASC
	`
	if err := r.db.SelectContext(ctx, &rows, query); err != nil {
		return nil, err
	}

	categories := make([]usecaseProductQuery.Classification, 0, len(rows))
	for _, row := range rows {
		categories = append(categories, usecaseProductQuery.Classification{
			UUID: row.UUID,
			Name: row.Name,
		})
	}

	return categories, nil
}

func (r *ProductQueryReader) findCategoryByUUID(ctx context.Context, categoryUUID string) (usecaseProductQuery.Classification, bool, error) {
	row := categoryRow{}
	if err := r.db.GetContext(
		ctx,
		&row,
		`SELECT c.uuid, c.name FROM category c WHERE c.uuid = ? AND c.deleted_at IS NULL LIMIT 1`,
		categoryUUID,
	); err != nil {
		if err == sql.ErrNoRows {
			return usecaseProductQuery.Classification{}, false, nil
		}
		return usecaseProductQuery.Classification{}, false, err
	}

	return usecaseProductQuery.Classification{
		UUID: row.UUID,
		Name: row.Name,
	}, true, nil
}

func (r *ProductQueryReader) listCarouselBaseRows(
	ctx context.Context,
	limit int,
	recommendOnly bool,
	excludeProductIDs []uint,
) ([]productBaseRow, error) {
	if limit <= 0 {
		return []productBaseRow{}, nil
	}

	query := `
		SELECT
			p.id,
			p.uuid,
			p.name,
			p.description,
			p.price,
			p.is_active,
			p.is_recommend,
			c.uuid AS category_uuid,
			c.name AS category_name,
			t.uuid AS target_uuid,
			t.name AS target_name
		FROM product p
		LEFT JOIN category c ON c.id = p.category_id AND c.deleted_at IS NULL
		LEFT JOIN target t ON t.id = p.target_id AND t.deleted_at IS NULL
		WHERE p.deleted_at IS NULL
		  AND p.is_active = 1
		  AND EXISTS (
		      SELECT 1
		      FROM product_image pi
		      WHERE pi.product_id = p.id AND pi.deleted_at IS NULL
		  )
	`

	args := make([]any, 0, 2)
	if recommendOnly {
		query += ` AND p.is_recommend = 1`
	}
	if len(excludeProductIDs) > 0 {
		query += ` AND p.id NOT IN (?)`
		args = append(args, excludeProductIDs)
	}
	query += ` ORDER BY p.created_at DESC, p.id DESC LIMIT ?`
	args = append(args, limit)

	if len(excludeProductIDs) > 0 {
		var err error
		query, args, err = sqlx.In(query, args...)
		if err != nil {
			return nil, err
		}
	}
	query = r.db.Rebind(query)

	rows := []productBaseRow{}
	if err := r.db.SelectContext(ctx, &rows, query, args...); err != nil {
		return nil, err
	}
	return rows, nil
}
