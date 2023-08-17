// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package enum

// AccessGrant represents the access grants a token or sshkey can have.
// Keep as int64 to allow for simpler+faster lookup of grants for a given token
// as we don't have to store an array field or need to do a join / 2nd db call.
// Multiple grants can be combined using the bit-wise or operation.
// ASSUMPTION: we don't need more than 63 grants!
//
// NOTE: A grant is always restricted by the principal permissions
//
// TODO: Beter name, access grant and permission might be to close in terminology?
type AccessGrant int64

const (
	// no grants - useless token.
	AccessGrantNone AccessGrant = 0

	// privacy related grants.
	AccessGrantPublic  AccessGrant = 1 << 0 // 1
	AccessGrantPrivate AccessGrant = 1 << 1 // 2

	// api related grants (spaces / repos, ...).
	AccessGrantAPICreate AccessGrant = 1 << 10 // 1024
	AccessGrantAPIView   AccessGrant = 1 << 11 // 2048
	AccessGrantAPIEdit   AccessGrant = 1 << 12 // 4096
	AccessGrantAPIDelete AccessGrant = 1 << 13 // 8192

	// code related grants.
	AccessGrantCodeRead  AccessGrant = 1 << 20 // 1048576
	AccessGrantCodeWrite AccessGrant = 1 << 21 // 2097152

	// grants everything - for user sessions.
	AccessGrantAll AccessGrant = 1<<63 - 1
)

// DoesGrantContain checks whether the grants contain all grants in the provided grant.
func (g AccessGrant) Contains(grants AccessGrant) bool {
	return g&grants == grants
}

// CombineGrants combines all grants into a single grant.
// Note: duplicates are ignored.
func CombineGrants(grants ...AccessGrant) AccessGrant {
	res := AccessGrantNone

	for _, grant := range grants {
		res |= grant
	}

	return res
}