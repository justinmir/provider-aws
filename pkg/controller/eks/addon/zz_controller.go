/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package addon

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/eks"
	svcsdk "github.com/aws/aws-sdk-go/service/eks"
	svcsdkapi "github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/eks/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an Addon resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Addon in AWS"
	errUpdate        = "cannot update Addon in AWS"
	errDescribe      = "failed to describe Addon"
	errDelete        = "failed to delete Addon"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, cr *svcapitypes.Addon) (managed.TypedExternalClient[*svcapitypes.Addon], error) {
	sess, err := connectaws.GetConfigV1(ctx, c.kube, cr, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, cr *svcapitypes.Addon) (managed.ExternalObservation, error) {
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeAddonInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeAddonWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateAddon(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	upToDate := true
	diff := ""
	if !meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		upToDate, diff, err = e.isUpToDate(ctx, cr, resp)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
		}
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, cr *svcapitypes.Addon) (managed.ExternalCreation, error) {
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateAddonInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateAddonWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.Addon.AddonArn != nil {
		cr.Status.AtProvider.AddonARN = resp.Addon.AddonArn
	} else {
		cr.Status.AtProvider.AddonARN = nil
	}
	if resp.Addon.AddonName != nil {
		cr.Spec.ForProvider.AddonName = resp.Addon.AddonName
	} else {
		cr.Spec.ForProvider.AddonName = nil
	}
	if resp.Addon.AddonVersion != nil {
		cr.Spec.ForProvider.AddonVersion = resp.Addon.AddonVersion
	} else {
		cr.Spec.ForProvider.AddonVersion = nil
	}
	if resp.Addon.ClusterName != nil {
		cr.Status.AtProvider.ClusterName = resp.Addon.ClusterName
	} else {
		cr.Status.AtProvider.ClusterName = nil
	}
	if resp.Addon.ConfigurationValues != nil {
		cr.Spec.ForProvider.ConfigurationValues = resp.Addon.ConfigurationValues
	} else {
		cr.Spec.ForProvider.ConfigurationValues = nil
	}
	if resp.Addon.CreatedAt != nil {
		cr.Status.AtProvider.CreatedAt = &metav1.Time{*resp.Addon.CreatedAt}
	} else {
		cr.Status.AtProvider.CreatedAt = nil
	}
	if resp.Addon.Health != nil {
		f6 := &svcapitypes.AddonHealth{}
		if resp.Addon.Health.Issues != nil {
			f6f0 := []*svcapitypes.AddonIssue{}
			for _, f6f0iter := range resp.Addon.Health.Issues {
				f6f0elem := &svcapitypes.AddonIssue{}
				if f6f0iter.Code != nil {
					f6f0elem.Code = f6f0iter.Code
				}
				if f6f0iter.Message != nil {
					f6f0elem.Message = f6f0iter.Message
				}
				if f6f0iter.ResourceIds != nil {
					f6f0elemf2 := []*string{}
					for _, f6f0elemf2iter := range f6f0iter.ResourceIds {
						var f6f0elemf2elem string
						f6f0elemf2elem = *f6f0elemf2iter
						f6f0elemf2 = append(f6f0elemf2, &f6f0elemf2elem)
					}
					f6f0elem.ResourceIDs = f6f0elemf2
				}
				f6f0 = append(f6f0, f6f0elem)
			}
			f6.Issues = f6f0
		}
		cr.Status.AtProvider.Health = f6
	} else {
		cr.Status.AtProvider.Health = nil
	}
	if resp.Addon.MarketplaceInformation != nil {
		f7 := &svcapitypes.MarketplaceInformation{}
		if resp.Addon.MarketplaceInformation.ProductId != nil {
			f7.ProductID = resp.Addon.MarketplaceInformation.ProductId
		}
		if resp.Addon.MarketplaceInformation.ProductUrl != nil {
			f7.ProductURL = resp.Addon.MarketplaceInformation.ProductUrl
		}
		cr.Status.AtProvider.MarketplaceInformation = f7
	} else {
		cr.Status.AtProvider.MarketplaceInformation = nil
	}
	if resp.Addon.ModifiedAt != nil {
		cr.Status.AtProvider.ModifiedAt = &metav1.Time{*resp.Addon.ModifiedAt}
	} else {
		cr.Status.AtProvider.ModifiedAt = nil
	}
	if resp.Addon.Owner != nil {
		cr.Status.AtProvider.Owner = resp.Addon.Owner
	} else {
		cr.Status.AtProvider.Owner = nil
	}
	if resp.Addon.Publisher != nil {
		cr.Status.AtProvider.Publisher = resp.Addon.Publisher
	} else {
		cr.Status.AtProvider.Publisher = nil
	}
	if resp.Addon.ServiceAccountRoleArn != nil {
		cr.Spec.ForProvider.ServiceAccountRoleARN = resp.Addon.ServiceAccountRoleArn
	} else {
		cr.Spec.ForProvider.ServiceAccountRoleARN = nil
	}
	if resp.Addon.Status != nil {
		cr.Status.AtProvider.Status = resp.Addon.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.Addon.Tags != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range resp.Addon.Tags {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		cr.Spec.ForProvider.Tags = f13
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, cr *svcapitypes.Addon) (managed.ExternalUpdate, error) {
	input := GenerateUpdateAddonInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateAddonWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, cr *svcapitypes.Addon) (managed.ExternalDelete, error) {
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteAddonInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return managed.ExternalDelete{}, errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return managed.ExternalDelete{}, nil
	}
	resp, err := e.client.DeleteAddonWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

func (e *external) Disconnect(ctx context.Context) error {
	// Unimplemented, required by newer versions of crossplane-runtime
	return nil
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EKSAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.EKSAPI
	preObserve     func(context.Context, *svcapitypes.Addon, *svcsdk.DescribeAddonInput) error
	postObserve    func(context.Context, *svcapitypes.Addon, *svcsdk.DescribeAddonOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.AddonParameters, *svcsdk.DescribeAddonOutput) error
	isUpToDate     func(context.Context, *svcapitypes.Addon, *svcsdk.DescribeAddonOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.Addon, *svcsdk.CreateAddonInput) error
	postCreate     func(context.Context, *svcapitypes.Addon, *svcsdk.CreateAddonOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Addon, *svcsdk.DeleteAddonInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Addon, *svcsdk.DeleteAddonOutput, error) (managed.ExternalDelete, error)
	preUpdate      func(context.Context, *svcapitypes.Addon, *svcsdk.UpdateAddonInput) error
	postUpdate     func(context.Context, *svcapitypes.Addon, *svcsdk.UpdateAddonOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Addon, *svcsdk.DescribeAddonInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Addon, _ *svcsdk.DescribeAddonOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.AddonParameters, *svcsdk.DescribeAddonOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.Addon, *svcsdk.DescribeAddonOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.Addon, *svcsdk.CreateAddonInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Addon, _ *svcsdk.CreateAddonOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Addon, *svcsdk.DeleteAddonInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Addon, _ *svcsdk.DeleteAddonOutput, err error) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, err
}
func nopPreUpdate(context.Context, *svcapitypes.Addon, *svcsdk.UpdateAddonInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Addon, _ *svcsdk.UpdateAddonOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
