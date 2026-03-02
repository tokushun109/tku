package di

import (
	"fmt"

	clockInfra "github.com/tokushun109/tku/backend/internal/infra/clock"
	"github.com/tokushun109/tku/backend/internal/infra/config"
	cryptoInfra "github.com/tokushun109/tku/backend/internal/infra/crypto"
	mailInfra "github.com/tokushun109/tku/backend/internal/infra/mail/sendgrid"
	creemaMarketplace "github.com/tokushun109/tku/backend/internal/infra/marketplace/creema"
	s3Storage "github.com/tokushun109/tku/backend/internal/infra/storage/s3"
	uuidInfra "github.com/tokushun109/tku/backend/internal/infra/uuid"
	usecase "github.com/tokushun109/tku/backend/internal/usecase"
	usecaseCategory "github.com/tokushun109/tku/backend/internal/usecase/category"
	usecaseContact "github.com/tokushun109/tku/backend/internal/usecase/contact"
	usecaseCreator "github.com/tokushun109/tku/backend/internal/usecase/creator"
	usecaseHealth "github.com/tokushun109/tku/backend/internal/usecase/health"
	usecaseProduct "github.com/tokushun109/tku/backend/internal/usecase/product"
	usecaseProductCommand "github.com/tokushun109/tku/backend/internal/usecase/product/command"
	usecaseProductQuery "github.com/tokushun109/tku/backend/internal/usecase/product/query"
	usecaseSalesSite "github.com/tokushun109/tku/backend/internal/usecase/sales_site"
	usecaseSession "github.com/tokushun109/tku/backend/internal/usecase/session"
	usecaseSkillMarket "github.com/tokushun109/tku/backend/internal/usecase/skill_market"
	usecaseSns "github.com/tokushun109/tku/backend/internal/usecase/sns"
	usecaseTag "github.com/tokushun109/tku/backend/internal/usecase/tag"
	usecaseTarget "github.com/tokushun109/tku/backend/internal/usecase/target"
	usecaseUser "github.com/tokushun109/tku/backend/internal/usecase/user"
)

type usecases struct {
	health      usecaseHealth.Usecase
	category    usecaseCategory.Usecase
	target      usecaseTarget.Usecase
	tag         usecaseTag.Usecase
	sns         usecaseSns.Usecase
	salesSite   usecaseSalesSite.Usecase
	skillMarket usecaseSkillMarket.Usecase
	creator     usecaseCreator.Usecase
	contact     usecaseContact.Usecase
	session     usecaseSession.Usecase
	user        usecaseUser.Usecase
	product     usecaseProduct.Usecase
}

func newUsecases(repos *repositories, qrs *queries, cfg *config.Config, txManager usecase.TxManager) (*usecases, error) {
	// 入力側の依存関係のチェック
	if err := requireStructFieldsNonNil("repositories", repos); err != nil {
		return nil, err
	}
	if err := requireStructFieldsNonNil("queries", qrs); err != nil {
		return nil, err
	}
	if err := requireNonNil("config", cfg); err != nil {
		return nil, err
	}
	if err := requireNonNil("txManager", txManager); err != nil {
		return nil, err
	}

	// 汎用的なユースケースの依存関係の構築
	uuidGen := uuidInfra.NewGenerator()
	if err := requireNonNil("uuidGenerator", uuidGen); err != nil {
		return nil, err
	}
	clock := clockInfra.NewClock()
	if err := requireNonNil("clock", clock); err != nil {
		return nil, err
	}
	passwordHasher := cryptoInfra.NewPasswordHasherSHA1()
	if err := requireNonNil("passwordHasher", passwordHasher); err != nil {
		return nil, err
	}
	mailer := mailInfra.NewMailer(cfg.Env, cfg.SendGridAPIKey)
	if err := requireNonNil("mailer", mailer); err != nil {
		return nil, err
	}
	storage, err := s3Storage.NewStorage(cfg.APIBucketName, cfg.S3UsePathStyle)
	if err != nil {
		return nil, fmt.Errorf("new creator logo storage: %w", err)
	}
	if err := requireNonNil("creatorStorage", storage); err != nil {
		return nil, err
	}
	productDuplicateSource := creemaMarketplace.NewScraper()
	if err := requireNonNil("productDuplicateSource", productDuplicateSource); err != nil {
		return nil, err
	}

	// ドメイン固有のユースケースの依存関係の構築
	contactNotifier := usecaseContact.NewContactNotifier(mailer, repos.user, cfg.ContactSupportEmail)
	if err := requireNonNil("contactNotifier", contactNotifier); err != nil {
		return nil, err
	}
	sessionUC := usecaseSession.New(repos.session, cfg.SessionTTL, clock)
	if err := requireNonNil("sessionUsecase", sessionUC); err != nil {
		return nil, err
	}
	productQueryUC := usecaseProductQuery.New(qrs.product, storage)
	if err := requireNonNil("productQueryUsecase", productQueryUC); err != nil {
		return nil, err
	}
	productCommandUC := usecaseProductCommand.New(
		repos.product,
		repos.productImage,
		repos.siteDetail,
		repos.category,
		repos.target,
		repos.tag,
		repos.salesSite,
		productDuplicateSource,
		storage,
		uuidGen,
		txManager,
	)
	if err := requireNonNil("productCommandUsecase", productCommandUC); err != nil {
		return nil, err
	}

	ucs := &usecases{
		health:      usecaseHealth.New(repos.health),
		category:    usecaseCategory.New(repos.category, uuidGen, txManager),
		target:      usecaseTarget.New(repos.target, uuidGen, txManager),
		tag:         usecaseTag.New(repos.tag, uuidGen, txManager),
		sns:         usecaseSns.New(repos.sns, uuidGen),
		salesSite:   usecaseSalesSite.New(repos.salesSite, uuidGen, txManager),
		skillMarket: usecaseSkillMarket.New(repos.skillMarket, uuidGen),
		creator:     usecaseCreator.New(repos.creator, storage, uuidGen),
		contact:     usecaseContact.New(repos.contact, contactNotifier, uuidGen),
		session:     sessionUC,
		user:        usecaseUser.New(repos.user, repos.session, sessionUC, passwordHasher, uuidGen, clock, txManager),
		product:     usecaseProduct.New(productQueryUC, productCommandUC),
	}

	// 出力側の依存関係のチェック
	if err := requireStructFieldsNonNil("usecases", ucs); err != nil {
		return nil, err
	}

	return ucs, nil
}
