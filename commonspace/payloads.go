package commonspace

import (
	aclrecordproto "github.com/anytypeio/any-sync/commonspace/object/acl/aclrecordproto"
	"github.com/anytypeio/any-sync/commonspace/object/keychain"
	"github.com/anytypeio/any-sync/commonspace/object/tree/objecttree"
	"github.com/anytypeio/any-sync/commonspace/spacestorage"
	"github.com/anytypeio/any-sync/commonspace/spacesyncproto"
	"github.com/anytypeio/any-sync/util/cidutil"
	"github.com/anytypeio/any-sync/util/keys/asymmetric/signingkey"
	"hash/fnv"
	"math/rand"
	"time"
)

const (
	SpaceReserved         = "any-sync.space"
	SpaceDerivationScheme = "derivation.standard"
)

func storagePayloadForSpaceCreate(payload SpaceCreatePayload) (storagePayload spacestorage.SpaceStorageCreatePayload, err error) {
	// unmarshalling signing and encryption keys
	identity, err := payload.SigningKey.GetPublic().Raw()
	if err != nil {
		return
	}
	encPubKey, err := payload.EncryptionKey.GetPublic().Raw()
	if err != nil {
		return
	}

	// preparing header and space id
	spaceHeaderSeed := make([]byte, 32)
	_, err = rand.Read(spaceHeaderSeed)
	if err != nil {
		return
	}
	header := &spacesyncproto.SpaceHeader{
		Identity:           identity,
		Timestamp:          time.Now().Unix(),
		SpaceType:          payload.SpaceType,
		SpaceHeaderPayload: payload.SpacePayload,
		ReplicationKey:     payload.ReplicationKey,
		Seed:               spaceHeaderSeed,
	}
	marshalled, err := header.Marshal()
	if err != nil {
		return
	}
	signature, err := payload.SigningKey.Sign(marshalled)
	if err != nil {
		return
	}
	rawHeader := &spacesyncproto.RawSpaceHeader{SpaceHeader: marshalled, Signature: signature}
	marshalled, err = rawHeader.Marshal()
	if err != nil {
		return
	}
	id, err := cidutil.NewCidFromBytes(marshalled)
	spaceId := NewSpaceId(id, payload.ReplicationKey)
	rawHeaderWithId := &spacesyncproto.RawSpaceHeaderWithId{
		RawHeader: marshalled,
		Id:        spaceId,
	}

	// encrypting read key
	hasher := fnv.New64()
	_, err = hasher.Write(payload.ReadKey)
	if err != nil {
		return
	}
	readKeyHash := hasher.Sum64()
	encReadKey, err := payload.EncryptionKey.GetPublic().Encrypt(payload.ReadKey)
	if err != nil {
		return
	}

	// preparing acl
	aclRoot := &aclrecordproto.AclRoot{
		Identity:           identity,
		EncryptionKey:      encPubKey,
		SpaceId:            spaceId,
		EncryptedReadKey:   encReadKey,
		CurrentReadKeyHash: readKeyHash,
		Timestamp:          time.Now().Unix(),
	}
	rawWithId, err := marshalAclRoot(aclRoot, payload.SigningKey)
	if err != nil {
		return
	}

	builder := objecttree.NewChangeBuilder(keychain.NewKeychain(), nil)
	spaceSettingsSeed := make([]byte, 32)
	_, err = rand.Read(spaceSettingsSeed)
	if err != nil {
		return
	}

	_, settingsRoot, err := builder.BuildRoot(objecttree.InitialContent{
		AclHeadId:  rawWithId.Id,
		Identity:   aclRoot.Identity,
		SigningKey: payload.SigningKey,
		SpaceId:    spaceId,
		Seed:       spaceSettingsSeed,
		ChangeType: SpaceReserved,
		Timestamp:  time.Now().Unix(),
	})
	if err != nil {
		return
	}

	// creating storage
	storagePayload = spacestorage.SpaceStorageCreatePayload{
		AclWithId:           rawWithId,
		SpaceHeaderWithId:   rawHeaderWithId,
		SpaceSettingsWithId: settingsRoot,
	}
	return
}

func storagePayloadForSpaceDerive(payload SpaceDerivePayload) (storagePayload spacestorage.SpaceStorageCreatePayload, err error) {
	// unmarshalling signing and encryption keys
	identity, err := payload.SigningKey.GetPublic().Raw()
	if err != nil {
		return
	}
	signPrivKey, err := payload.SigningKey.Raw()
	if err != nil {
		return
	}
	encPubKey, err := payload.EncryptionKey.GetPublic().Raw()
	if err != nil {
		return
	}
	encPrivKey, err := payload.EncryptionKey.Raw()
	if err != nil {
		return
	}

	// preparing replication key
	hasher := fnv.New64()
	_, err = hasher.Write(identity)
	if err != nil {
		return
	}
	repKey := hasher.Sum64()

	// preparing header and space id
	header := &spacesyncproto.SpaceHeader{
		Identity:           identity,
		SpaceType:          payload.SpaceType,
		SpaceHeaderPayload: payload.SpacePayload,
		ReplicationKey:     repKey,
	}
	marshalled, err := header.Marshal()
	if err != nil {
		return
	}
	signature, err := payload.SigningKey.Sign(marshalled)
	if err != nil {
		return
	}
	rawHeader := &spacesyncproto.RawSpaceHeader{SpaceHeader: marshalled, Signature: signature}
	marshalled, err = rawHeader.Marshal()
	if err != nil {
		return
	}
	id, err := cidutil.NewCidFromBytes(marshalled)
	spaceId := NewSpaceId(id, repKey)
	rawHeaderWithId := &spacesyncproto.RawSpaceHeaderWithId{
		RawHeader: marshalled,
		Id:        spaceId,
	}

	// deriving and encrypting read key
	readKey, err := aclrecordproto.AclReadKeyDerive(signPrivKey, encPrivKey)
	if err != nil {
		return
	}
	hasher = fnv.New64()
	_, err = hasher.Write(readKey.Bytes())
	if err != nil {
		return
	}
	readKeyHash := hasher.Sum64()

	// preparing acl
	aclRoot := &aclrecordproto.AclRoot{
		Identity:           identity,
		EncryptionKey:      encPubKey,
		SpaceId:            spaceId,
		DerivationScheme:   SpaceDerivationScheme,
		CurrentReadKeyHash: readKeyHash,
	}
	rawWithId, err := marshalAclRoot(aclRoot, payload.SigningKey)
	if err != nil {
		return
	}

	builder := objecttree.NewChangeBuilder(keychain.NewKeychain(), nil)
	_, settingsRoot, err := builder.BuildRoot(objecttree.InitialContent{
		AclHeadId:  rawWithId.Id,
		Identity:   aclRoot.Identity,
		SigningKey: payload.SigningKey,
		SpaceId:    spaceId,
		ChangeType: SpaceReserved,
	})
	if err != nil {
		return
	}

	// creating storage
	storagePayload = spacestorage.SpaceStorageCreatePayload{
		AclWithId:           rawWithId,
		SpaceHeaderWithId:   rawHeaderWithId,
		SpaceSettingsWithId: settingsRoot,
	}
	return
}

func marshalAclRoot(aclRoot *aclrecordproto.AclRoot, key signingkey.PrivKey) (rawWithId *aclrecordproto.RawAclRecordWithId, err error) {
	marshalledRoot, err := aclRoot.Marshal()
	if err != nil {
		return
	}
	signature, err := key.Sign(marshalledRoot)
	if err != nil {
		return
	}
	raw := &aclrecordproto.RawAclRecord{
		Payload:   marshalledRoot,
		Signature: signature,
	}
	marshalledRaw, err := raw.Marshal()
	if err != nil {
		return
	}
	aclHeadId, err := cidutil.NewCidFromBytes(marshalledRaw)
	if err != nil {
		return
	}
	rawWithId = &aclrecordproto.RawAclRecordWithId{
		Payload: marshalledRaw,
		Id:      aclHeadId,
	}
	return
}
