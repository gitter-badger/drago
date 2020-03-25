package server

import (
	gomodel "gopkg.in/jeevatkm/go-model.v1"
)

type Controller interface {
	CreateHost(*CreateHostInput) (*Host, error)
	UpdateHost(*UpdateHostInput) (*Host, error)
	DeleteHost(*DeleteHostInput) error
	GetAllHosts(*GetAllHostsInput) ([]*Host, error)
	GetHost(*GetHostInput) (*Host, error)
	SyncHost(*SyncHostInput) (*Host, error)
	CreateLink(*CreateLinkInput) (*Link, error)
	UpdateLink(*UpdateLinkInput) (*Link, error)
	GetAllLinks(*GetAllLinksInput) ([]*Link, error)
	DeleteLink(*DeleteLinkInput) error
}

type controller struct {
	repo Repository
}

func NewController(r Repository) (Controller, error) {
	return &controller{
		repo: r,
	}, nil
}

func (c *controller) CreateHost(i *CreateHostInput) (*Host, error) {
	h := &Host{}
	gomodel.Copy(h, i)

	hh, err := c.repo.CreateHost(h)

	return hh, err
}

func (c *controller) UpdateHost(i *UpdateHostInput) (*Host, error) {
	h := &Host{}
	gomodel.Copy(h, i)
	return c.repo.UpdateHost(i.ID, h)
}

func (c *controller) DeleteHost(i *DeleteHostInput) error {
	return c.repo.DeleteHost(i.ID)
}

func (c *controller) GetHost(i *GetHostInput) (*Host, error) {
	return c.repo.GetHost(i.ID)
}

func (c *controller) SyncHost(i *SyncHostInput) (*Host, error) {
	c.repo.UpdateHost(i.ID, &Host{
		ID:        i.ID,
		PublicKey: i.PublicKey,
	})
	return c.repo.GetHost(i.ID)
}

func (c *controller) GetAllHosts(i *GetAllHostsInput) ([]*Host, error) {
	return c.repo.GetAllHosts()
}

func (c *controller) GetAllLinks(i *GetAllLinksInput) ([]*Link, error) {
	return c.repo.GetAllLinks()
}

func (c *controller) GetAllLinksFromNode(id int) ([]*Link, error) {
	return nil, nil
}

func (c *controller) CreateLink(i *CreateLinkInput) (*Link, error) {
	l := Link{}
	gomodel.Copy(&l, i)
	return c.repo.CreateLink(&l)
}

func (c *controller) UpdateLink(i *UpdateLinkInput) (*Link, error) {
	l := &Link{}
	gomodel.Copy(&l, i)
	return c.repo.UpdateLink(i.ID, l)
}

func (c *controller) DeleteLink(i *DeleteLinkInput) error {
	return c.repo.DeleteLink(i.ID)
}