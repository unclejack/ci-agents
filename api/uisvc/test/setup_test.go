package test

import (
	"testing"
	"time"

	check "github.com/erikh/check"
	"github.com/tinyci/ci-agents/clients/asset"
	"github.com/tinyci/ci-agents/grpc/handler"
	"github.com/tinyci/ci-agents/testutil"
	"github.com/tinyci/ci-agents/testutil/testclients"
	"github.com/tinyci/ci-agents/testutil/testservers"
)

type uisvcSuite struct {
	datasvcClient  *testclients.DataClient
	queuesvcClient *testclients.QueueClient
	assetsvcClient *asset.Client
	logDoneChan    chan struct{}
	queueDoneChan  chan struct{}
	dataDoneChan   chan struct{}
	assetDoneChan  chan struct{}
	logHandler     *handler.H
	dataHandler    *handler.H
	queueHandler   *handler.H
	assetHandler   *handler.H

	logJournal *testservers.LogJournal
}

var _ = check.Suite(&uisvcSuite{})

func TestUISvc(t *testing.T) {
	check.TestingT(t)
}

func (us *uisvcSuite) SetUpTest(c *check.C) {
	testutil.WipeDB(c)

	var err error
	us.dataHandler, us.dataDoneChan, err = testservers.MakeDataServer()
	c.Assert(err, check.IsNil)

	us.queueHandler, us.queueDoneChan, err = testservers.MakeQueueServer()
	c.Assert(err, check.IsNil)

	us.assetHandler, us.assetDoneChan, err = testservers.MakeAssetServer()
	c.Assert(err, check.IsNil)

	us.logHandler, us.logDoneChan, us.logJournal, err = testservers.MakeLogServer()
	c.Assert(err, check.IsNil)

	go us.logJournal.Tail()

	us.datasvcClient, err = testclients.NewDataClient()
	c.Assert(err, check.IsNil)

	us.queuesvcClient, err = testclients.NewQueueClient(us.datasvcClient)
	c.Assert(err, check.IsNil)

	us.assetsvcClient, err = asset.NewClient(nil, "localhost:6002")
	c.Assert(err, check.IsNil)
}

func (us *uisvcSuite) TearDownTest(c *check.C) {
	close(us.dataDoneChan)
	close(us.queueDoneChan)
	close(us.logDoneChan)
	close(us.assetDoneChan)
	time.Sleep(100 * time.Millisecond)
}
