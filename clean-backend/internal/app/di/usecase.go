package di

import (
	"fmt"
	"strings"

	clockInfra "github.com/tokushun109/tku/clean-backend/internal/infra/clock"
	"github.com/tokushun109/tku/clean-backend/internal/infra/config"
	cryptoInfra "github.com/tokushun109/tku/clean-backend/internal/infra/crypto"
	mailInfra "github.com/tokushun109/tku/clean-backend/internal/infra/mail/sendgrid"
	localStorage "github.com/tokushun109/tku/clean-backend/internal/infra/storage/local"
	s3Storage "github.com/tokushun109/tku/clean-backend/internal/infra/storage/s3"
	uuidInfra "github.com/tokushun109/tku/clean-backend/internal/infra/uuid"
	usecase "github.com/tokushun109/tku/clean-backend/internal/usecase"
	usecaseCategory "github.com/tokushun109/tku/clean-backend/internal/usecase/category"
	usecaseContact "github.com/tokushun109/tku/clean-backend/internal/usecase/contact"
	usecaseCreator "github.com/tokushun109/tku/clean-backend/internal/usecase/creator"
	usecaseHealth "github.com/tokushun109/tku/clean-backend/internal/usecase/health"
	usecaseSalesSite "github.com/tokushun109/tku/clean-backend/internal/usecase/sales_site"
	usecaseSession "github.com/tokushun109/tku/clean-backend/internal/usecase/session"
	usecaseSkillMarket "github.com/tokushun109/tku/clean-backend/internal/usecase/skill_market"
	usecaseSns "github.com/tokushun109/tku/clean-backend/internal/usecase/sns"
	usecaseTag "github.com/tokushun109/tku/clean-backend/internal/usecase/tag"
	usecaseTarget "github.com/tokushun109/tku/clean-backend/internal/usecase/target"
	usecaseUser "github.com/tokushun109/tku/clean-backend/internal/usecase/user"
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
}

func newUsecases(repos *repositories, cfg *config.Config, txManager usecase.TxManager) (*usecases, error) {
	// 入力側の依存関係のチェック
	if err := requireStructFieldsNonNil("repositories", repos); err != nil {
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
	storage, err := newLogoStorage(cfg)
	if err != nil {
		return nil, fmt.Errorf("new creator logo storage: %w", err)
	}
	if err := requireNonNil("creatorStorage", storage); err != nil {
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

	ucs := &usecases{
		health:      usecaseHealth.New(repos.health),
		category:    usecaseCategory.New(repos.category, uuidGen),
		target:      usecaseTarget.New(repos.target, uuidGen),
		tag:         usecaseTag.New(repos.tag, uuidGen),
		sns:         usecaseSns.New(repos.sns, uuidGen),
		salesSite:   usecaseSalesSite.New(repos.salesSite, uuidGen),
		skillMarket: usecaseSkillMarket.New(repos.skillMarket, uuidGen),
		creator:     usecaseCreator.New(repos.creator, storage, uuidGen, cfg.Env, cfg.APIBaseURL, 0),
		contact:     usecaseContact.New(repos.contact, contactNotifier),
		session:     sessionUC,
		user:        usecaseUser.New(repos.user, repos.session, sessionUC, passwordHasher, uuidGen, clock, txManager),
	}

	// 出力側の依存関係のチェック
	if err := requireStructFieldsNonNil("usecases", ucs); err != nil {
		return nil, err
	}

	return ucs, nil
}

func newLogoStorage(cfg *config.Config) (usecase.Storage, error) {
	if strings.EqualFold(strings.TrimSpace(cfg.Env), "local") || strings.TrimSpace(cfg.Env) == "" {
		return localStorage.NewStorage("."), nil
	}
	return s3Storage.NewStorage(cfg.APIBucketName)
}
