package telegram_test

import (
	"context"
	"crypto/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"golang.org/x/xerrors"

	"github.com/gotd/td/internal/tdsync"
	"github.com/gotd/td/internal/testutil"
	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/internal/e2etest"
	"github.com/gotd/td/telegram/internal/tgtest"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/transport"
)

func testTransport(trp telegram.Transport, storage session.Storage) func(t *testing.T) {
	return func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		log := zaptest.NewLogger(t)
		defer func() { _ = log.Sync() }()

		err := telegram.TestClient(ctx, telegram.Options{
			Transport:      trp,
			Logger:         log.Named("client"),
			SessionStorage: storage,
		}, func(ctx context.Context, client *telegram.Client) error {
			if _, err := client.Self(ctx); err != nil {
				return xerrors.Errorf("self: %w", err)
			}

			return nil
		})

		require.NoError(t, err)
	}
}

func TestExternalE2EConnect(t *testing.T) {
	testutil.SkipExternal(t)

	// To re-use session.
	storage := &session.StorageMemory{}
	t.Run("Abridged", testTransport(transport.Abridged(nil), storage))
	t.Run("Intermediate", testTransport(transport.Intermediate(nil), storage))
	t.Run("PaddedIntermediate", testTransport(transport.PaddedIntermediate(nil), storage))
	t.Run("Full", testTransport(transport.Full(nil), storage))
}

const dialog = `— Да?
— Алё!
— Да да?
— Ну как там с деньгами?
— А?
— Как с деньгами-то там?
— Чё с деньгами?
— Чё?
— Куда ты звонишь?
— Тебе звоню.
— Кому?
— Ну тебе.`

func TestExternalE2EUsersDialog(t *testing.T) {
	testutil.SkipExternal(t)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	log := zaptest.NewLogger(t).WithOptions(zap.IncreaseLevel(zapcore.InfoLevel))

	cfg := e2etest.TestConfig{
		AppID:   telegram.TestAppID,
		AppHash: telegram.TestAppHash,
		DcID:    2,
		Addr:    telegram.AddrTest,
	}
	suite := e2etest.NewSuite(tgtest.NewSuite(ctx, t, log), cfg, rand.Reader)

	auth := make(chan *tg.User, 1)
	g := tdsync.NewLogGroup(ctx, log.Named("group"))

	g.Go("echobot", func(ctx context.Context) error {
		return e2etest.NewEchoBot(suite, auth).Run(ctx)
	})

	user, ok := <-auth
	if ok {
		g.Go("terentyev", func(ctx context.Context) error {
			defer g.Cancel()
			return e2etest.NewUser(suite, strings.Split(dialog, "\n"), user.Username).Run(ctx)
		})
	}

	require.NoError(t, g.Wait())
}
