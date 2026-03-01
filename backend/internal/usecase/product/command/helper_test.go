package command

import (
	"context"
	"errors"
	"testing"

	domainCategory "github.com/tokushun109/tku/backend/internal/domain/category"
	"github.com/tokushun109/tku/backend/internal/domain/primitive"
	domainSalesSite "github.com/tokushun109/tku/backend/internal/domain/sales_site"
	domainTag "github.com/tokushun109/tku/backend/internal/domain/tag"
	domainTarget "github.com/tokushun109/tku/backend/internal/domain/target"
	"github.com/tokushun109/tku/backend/internal/usecase"
	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
)

const (
	helperCategoryUUID  = "11111111-1111-4111-8111-111111111111"
	helperTargetUUID    = "22222222-2222-4222-8222-222222222222"
	helperTagUUID       = "33333333-3333-4333-8333-333333333333"
	helperSalesSiteUUID = "44444444-4444-4444-8444-444444444444"
	helperProductUUID   = "55555555-5555-4555-8555-555555555555"
	helperCreatedUUID   = "66666666-6666-4666-8666-666666666666"
)

type helperCategoryRepoStub struct {
	byUUID map[string]*domainCategory.Category
}

func (s *helperCategoryRepoStub) Create(ctx context.Context, c *domainCategory.Category) (*domainCategory.Category, error) {
	return c, nil
}

func (s *helperCategoryRepoStub) FindAll(ctx context.Context) ([]*domainCategory.Category, error) {
	return nil, nil
}

func (s *helperCategoryRepoStub) FindUsed(ctx context.Context) ([]*domainCategory.Category, error) {
	return nil, nil
}

func (s *helperCategoryRepoStub) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainCategory.Category, error) {
	if s.byUUID == nil {
		return nil, nil
	}
	return s.byUUID[uuid.Value()], nil
}

func (s *helperCategoryRepoStub) FindByName(ctx context.Context, name domainCategory.CategoryName) (*domainCategory.Category, error) {
	return nil, nil
}

func (s *helperCategoryRepoStub) ExistsByName(ctx context.Context, name domainCategory.CategoryName) (bool, error) {
	return false, nil
}

func (s *helperCategoryRepoStub) Update(ctx context.Context, c *domainCategory.Category) (bool, error) {
	return false, nil
}

func (s *helperCategoryRepoStub) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

type helperTargetRepoStub struct {
	byUUID map[string]*domainTarget.Target
}

func (s *helperTargetRepoStub) Create(ctx context.Context, t *domainTarget.Target) (*domainTarget.Target, error) {
	return t, nil
}

func (s *helperTargetRepoStub) FindAll(ctx context.Context) ([]*domainTarget.Target, error) {
	return nil, nil
}

func (s *helperTargetRepoStub) FindUsed(ctx context.Context) ([]*domainTarget.Target, error) {
	return nil, nil
}

func (s *helperTargetRepoStub) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainTarget.Target, error) {
	if s.byUUID == nil {
		return nil, nil
	}
	return s.byUUID[uuid.Value()], nil
}

func (s *helperTargetRepoStub) FindByName(ctx context.Context, name domainTarget.TargetName) (*domainTarget.Target, error) {
	return nil, nil
}

func (s *helperTargetRepoStub) ExistsByName(ctx context.Context, name domainTarget.TargetName) (bool, error) {
	return false, nil
}

func (s *helperTargetRepoStub) Update(ctx context.Context, t *domainTarget.Target) (bool, error) {
	return false, nil
}

func (s *helperTargetRepoStub) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

type helperTagRepoStub struct {
	byUUID            map[string]*domainTag.Tag
	findByUUIDsCalled int
}

func (s *helperTagRepoStub) Create(ctx context.Context, t *domainTag.Tag) (*domainTag.Tag, error) {
	return t, nil
}

func (s *helperTagRepoStub) FindAll(ctx context.Context) ([]*domainTag.Tag, error) {
	return nil, nil
}

func (s *helperTagRepoStub) FindByName(ctx context.Context, name domainTag.TagName) (*domainTag.Tag, error) {
	return nil, nil
}

func (s *helperTagRepoStub) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainTag.Tag, error) {
	if s.byUUID == nil {
		return nil, nil
	}
	return s.byUUID[uuid.Value()], nil
}

func (s *helperTagRepoStub) FindByUUIDs(ctx context.Context, uuids []primitive.UUID) ([]*domainTag.Tag, error) {
	s.findByUUIDsCalled++
	if s.byUUID == nil {
		return []*domainTag.Tag{}, nil
	}

	tags := make([]*domainTag.Tag, 0, len(uuids))
	for _, uuid := range uuids {
		tag, ok := s.byUUID[uuid.Value()]
		if !ok {
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (s *helperTagRepoStub) ExistsByName(ctx context.Context, name domainTag.TagName) (bool, error) {
	return false, nil
}

func (s *helperTagRepoStub) Update(ctx context.Context, t *domainTag.Tag) (bool, error) {
	return false, nil
}

func (s *helperTagRepoStub) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

type helperSalesSiteRepoStub struct {
	byUUID map[string]*domainSalesSite.SalesSite
}

func (s *helperSalesSiteRepoStub) Create(ctx context.Context, salesSite *domainSalesSite.SalesSite) (*domainSalesSite.SalesSite, error) {
	return salesSite, nil
}

func (s *helperSalesSiteRepoStub) FindAll(ctx context.Context) ([]*domainSalesSite.SalesSite, error) {
	return nil, nil
}

func (s *helperSalesSiteRepoStub) FindByName(ctx context.Context, name domainSalesSite.SalesSiteName) (*domainSalesSite.SalesSite, error) {
	return nil, nil
}

func (s *helperSalesSiteRepoStub) FindByUUID(ctx context.Context, uuid primitive.UUID) (*domainSalesSite.SalesSite, error) {
	if s.byUUID == nil {
		return nil, nil
	}
	return s.byUUID[uuid.Value()], nil
}

func (s *helperSalesSiteRepoStub) Update(ctx context.Context, salesSite *domainSalesSite.SalesSite) (bool, error) {
	return false, nil
}

func (s *helperSalesSiteRepoStub) Delete(ctx context.Context, uuid primitive.UUID) (bool, error) {
	return false, nil
}

type helperUUIDGenStub struct{}

func (g *helperUUIDGenStub) New() string {
	return helperCreatedUUID
}

func TestResolveCategoryUUIDReturnsInvalidInputWhenCategoryDoesNotExist(t *testing.T) {
	s := &Service{
		categoryRepo: &helperCategoryRepoStub{},
	}

	_, err := s.resolveCategoryUUID(context.Background(), helperCategoryUUID)
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestResolveTargetUUIDReturnsInvalidInputWhenTargetDoesNotExist(t *testing.T) {
	s := &Service{
		targetRepo: &helperTargetRepoStub{},
	}

	_, err := s.resolveTargetUUID(context.Background(), helperTargetUUID)
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestResolveTagUUIDsReturnsInvalidInputWhenTagDoesNotExist(t *testing.T) {
	repo := &helperTagRepoStub{}
	s := &Service{tagRepo: repo}

	_, err := s.resolveTagUUIDs(context.Background(), []string{helperTagUUID})
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
	if repo.findByUUIDsCalled != 1 {
		t.Fatalf("expected FindByUUIDs called once, got %d", repo.findByUUIDsCalled)
	}
}

func TestBuildSiteDetailsReturnsInvalidInputWhenSalesSiteDoesNotExist(t *testing.T) {
	productUUID, err := primitive.NewUUID(helperProductUUID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	s := &Service{
		salesSiteRepo: &helperSalesSiteRepoStub{},
		uuidGen:       &helperUUIDGenStub{},
	}

	_, err = s.buildSiteDetails(context.Background(), productUUID, []usecaseProduct.SiteDetailInput{
		{
			DetailURL:     "https://example.com/items/1",
			SalesSiteUUID: helperSalesSiteUUID,
		},
	})
	if err == nil || !errors.Is(err, usecase.ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}
