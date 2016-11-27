package srrlinepub

import "github.com/Sirupsen/logrus"

// Ctx ...
type Ctx struct {
	*config
	*logger
}

// NewCtx ...
func NewCtx(configpath string) (*Ctx, error) {
	config, cerr := newConfig(configpath)
	if cerr != nil {
		logrus.Errorf("[ctx][NewCtx] %+v\n", cerr)
		return nil, cerr
	}

	logger, lerr := newLogger(config)
	if lerr != nil {
		logrus.Errorf("[ctx][NewCtx] %+v\n", lerr)
		return nil, lerr
	}

	return &Ctx{config: config, logger: logger}, nil
}

// Close ...
func (c *Ctx) Close() error {
	c.logger.close()
	return nil
}
