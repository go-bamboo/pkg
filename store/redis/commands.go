package redis

import "github.com/go-redis/redis/v8"

//import (
//	"context"
//	"time"
//
//	"github.com/go-redis/redis/v8"
//)
//
//type (
//	Pipeliner              = redis.Pipeline
//	Cmder                  = redis.Cmder
//	CommandsInfoCmd        = redis.CommandsInfoCmd
//	StringCmd              = redis.StringCmd
//	StatusCmd              = redis.StatusCmd
//	IntCmd                 = redis.IntCmd
//	BoolCmd                = redis.BoolCmd
//	StringSliceCmd         = redis.StringSliceCmd
//	DurationCmd            = redis.DurationCmd
//	Sort                   = redis.Sort
//	GeoLocation            = redis.GeoLocation
//	SliceCmd               = redis.SliceCmd
//	FloatCmd               = redis.FloatCmd
//	ScanCmd                = redis.ScanCmd
//	BitCount               = redis.BitCount
//	IntSliceCmd            = redis.IntSliceCmd
//	LPosArgs               = redis.LPosArgs
//	XClaimArgs             = redis.XClaimArgs
//	XMessageSliceCmd       = redis.XMessageSliceCmd
//	XStreamSliceCmd        = redis.XStreamSliceCmd
//	XAddArgs               = redis.XAddArgs
//	XInfoGroupsCmd         = redis.XInfoGroupsCmd
//	StringStringMapCmd     = redis.StringStringMapCmd
//	GeoPosCmd              = redis.GeoPosCmd
//	SetArgs                = redis.SetArgs
//	BoolSliceCmd           = redis.BoolSliceCmd
//	XReadGroupArgs         = redis.XReadGroupArgs
//	XReadArgs              = redis.XReadArgs
//	XPendingCmd            = redis.XPendingCmd
//	XPendingExtArgs        = redis.XPendingExtArgs
//	FloatSliceCmd          = redis.FloatSliceCmd
//	XAutoClaimJustIDCmd    = redis.XAutoClaimJustIDCmd
//	XPendingExtCmd         = redis.XPendingExtCmd
//	XAutoClaimCmd          = redis.XAutoClaimCmd
//	XInfoStreamCmd         = redis.XInfoStreamCmd
//	StringStructMapCmd     = redis.StringStructMapCmd
//	XAutoClaimArgs         = redis.XAutoClaimArgs
//	XInfoStreamFullCmd     = redis.XInfoStreamFullCmd
//	XInfoConsumersCmd      = redis.XInfoConsumersCmd
//	ZWithKeyCmd            = redis.ZWithKeyCmd
//	TimeCmd                = redis.TimeCmd
//	Z                      = redis.Z
//	ZSliceCmd              = redis.ZSliceCmd
//	ZAddArgs               = redis.ZAddArgs
//	ZStore                 = redis.ZStore
//	ZRangeBy               = redis.ZRangeBy
//	Cmd                    = redis.Cmd
//	ZRangeArgs             = redis.ZRangeArgs
//	StringIntMapCmd        = redis.StringIntMapCmd
//	ClusterSlotsCmd        = redis.ClusterSlotsCmd
//	GeoRadiusQuery         = redis.GeoRadiusQuery
//	GeoLocationCmd         = redis.GeoLocationCmd
//	GeoSearchLocationCmd   = redis.GeoSearchLocationCmd
//	GeoSearchQuery         = redis.GeoSearchQuery
//	GeoSearchLocationQuery = redis.GeoSearchLocationQuery
//	GeoSearchStoreQuery    = redis.GeoSearchStoreQuery
//)
//
//type Client interface {
//	//Auth(ctx context.Context, password string) *StatusCmd
//	//AuthACL(ctx context.Context, username, password string) *StatusCmd
//	//Select(ctx context.Context, index int) *StatusCmd
//	//SwapDB(ctx context.Context, index1, index2 int) *StatusCmd
//	//ClientSetName(ctx context.Context, name string) *BoolCmd
//
//	Pipeline() Pipeliner
//	Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)
//
//	TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)
//	TxPipeline() Pipeliner
//
//	Command(ctx context.Context) *CommandsInfoCmd
//	ClientGetName(ctx context.Context) *StringCmd
//	Echo(ctx context.Context, message interface{}) *StringCmd
//	Ping(ctx context.Context) *StatusCmd
//	Quit(ctx context.Context) *StatusCmd
//	Del(ctx context.Context, keys ...string) *IntCmd
//	Unlink(ctx context.Context, keys ...string) *IntCmd
//	Dump(ctx context.Context, key string) *StringCmd
//	Exists(ctx context.Context, keys ...string) *IntCmd
//	Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
//	ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
//	Keys(ctx context.Context, pattern string) *StringSliceCmd
//	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *StatusCmd
//	Move(ctx context.Context, key string, db int) *BoolCmd
//	ObjectRefCount(ctx context.Context, key string) *IntCmd
//	ObjectEncoding(ctx context.Context, key string) *StringCmd
//	ObjectIdleTime(ctx context.Context, key string) *DurationCmd
//	Persist(ctx context.Context, key string) *BoolCmd
//	PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
//	PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
//	PTTL(ctx context.Context, key string) *DurationCmd
//	RandomKey(ctx context.Context) *StringCmd
//	Rename(ctx context.Context, key, newkey string) *StatusCmd
//	RenameNX(ctx context.Context, key, newkey string) *BoolCmd
//	Restore(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
//	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
//	Sort(ctx context.Context, key string, sort *Sort) *StringSliceCmd
//	SortStore(ctx context.Context, key, store string, sort *Sort) *IntCmd
//	SortInterfaces(ctx context.Context, key string, sort *Sort) *SliceCmd
//	Touch(ctx context.Context, keys ...string) *IntCmd
//	TTL(ctx context.Context, key string) *DurationCmd
//	Type(ctx context.Context, key string) *StatusCmd
//	Append(ctx context.Context, key, value string) *IntCmd
//	Decr(ctx context.Context, key string) *IntCmd
//	DecrBy(ctx context.Context, key string, decrement int64) *IntCmd
//	Get(ctx context.Context, key string) *StringCmd
//	GetRange(ctx context.Context, key string, start, end int64) *StringCmd
//	GetSet(ctx context.Context, key string, value interface{}) *StringCmd
//	GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd
//	GetDel(ctx context.Context, key string) *StringCmd
//	Incr(ctx context.Context, key string) *IntCmd
//	IncrBy(ctx context.Context, key string, value int64) *IntCmd
//	IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd
//	MGet(ctx context.Context, keys ...string) *SliceCmd
//	MSet(ctx context.Context, values ...interface{}) *StatusCmd
//	MSetNX(ctx context.Context, values ...interface{}) *BoolCmd
//	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
//	SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *StatusCmd
//	// TODO: rename to SetEx
//	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
//	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
//	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
//	SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
//	StrLen(ctx context.Context, key string) *IntCmd
//
//	GetBit(ctx context.Context, key string, offset int64) *IntCmd
//	SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd
//	BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
//	BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd
//	BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd
//	BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd
//	BitOpNot(ctx context.Context, destKey string, key string) *IntCmd
//	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd
//	BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd
//
//	Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
//	ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd
//	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
//	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
//	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
//
//	HDel(ctx context.Context, key string, fields ...string) *IntCmd
//	HExists(ctx context.Context, key, field string) *BoolCmd
//	HGet(ctx context.Context, key, field string) *StringCmd
//	HGetAll(ctx context.Context, key string) *StringStringMapCmd
//	HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd
//	HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd
//	HKeys(ctx context.Context, key string) *StringSliceCmd
//	HLen(ctx context.Context, key string) *IntCmd
//	HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
//	HSet(ctx context.Context, key string, values ...interface{}) *IntCmd
//	HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd
//	HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd
//	HVals(ctx context.Context, key string) *StringSliceCmd
//	HRandField(ctx context.Context, key string, count int, withValues bool) *StringSliceCmd
//
//	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
//	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
//	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd
//	LIndex(ctx context.Context, key string, index int64) *StringCmd
//	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd
//	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd
//	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd
//	LLen(ctx context.Context, key string) *IntCmd
//	LPop(ctx context.Context, key string) *StringCmd
//	LPopCount(ctx context.Context, key string, count int) *StringSliceCmd
//	LPos(ctx context.Context, key string, value string, args LPosArgs) *IntCmd
//	LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd
//	LPush(ctx context.Context, key string, values ...interface{}) *IntCmd
//	LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
//	LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
//	LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd
//	LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd
//	LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
//	RPop(ctx context.Context, key string) *StringCmd
//	RPopCount(ctx context.Context, key string, count int) *StringSliceCmd
//	RPopLPush(ctx context.Context, source, destination string) *StringCmd
//	RPush(ctx context.Context, key string, values ...interface{}) *IntCmd
//	RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
//	LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd
//	BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd
//
//	SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd
//	SCard(ctx context.Context, key string) *IntCmd
//	SDiff(ctx context.Context, keys ...string) *StringSliceCmd
//	SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
//	SInter(ctx context.Context, keys ...string) *StringSliceCmd
//	SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd
//	SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd
//	SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd
//	SMembers(ctx context.Context, key string) *StringSliceCmd
//	SMembersMap(ctx context.Context, key string) *StringStructMapCmd
//	SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd
//	SPop(ctx context.Context, key string) *StringCmd
//	SPopN(ctx context.Context, key string, count int64) *StringSliceCmd
//	SRandMember(ctx context.Context, key string) *StringCmd
//	SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
//	SRem(ctx context.Context, key string, members ...interface{}) *IntCmd
//	SUnion(ctx context.Context, keys ...string) *StringSliceCmd
//	SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd
//
//	XAdd(ctx context.Context, a *XAddArgs) *StringCmd
//	XDel(ctx context.Context, stream string, ids ...string) *IntCmd
//	XLen(ctx context.Context, stream string) *IntCmd
//	XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
//	XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
//	XRevRange(ctx context.Context, stream string, start, stop string) *XMessageSliceCmd
//	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd
//	XRead(ctx context.Context, a *XReadArgs) *XStreamSliceCmd
//	XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd
//	XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd
//	XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd
//	XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd
//	XGroupDestroy(ctx context.Context, stream, group string) *IntCmd
//	XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
//	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
//	XReadGroup(ctx context.Context, a *XReadGroupArgs) *XStreamSliceCmd
//	XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd
//	XPending(ctx context.Context, stream, group string) *XPendingCmd
//	XPendingExt(ctx context.Context, a *XPendingExtArgs) *XPendingExtCmd
//	XClaim(ctx context.Context, a *XClaimArgs) *XMessageSliceCmd
//	XClaimJustID(ctx context.Context, a *XClaimArgs) *StringSliceCmd
//	XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimCmd
//	XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimJustIDCmd
//
//	// TODO: XTrim and XTrimApprox remove in v9.
//	XTrim(ctx context.Context, key string, maxLen int64) *IntCmd
//	XTrimApprox(ctx context.Context, key string, maxLen int64) *IntCmd
//	XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd
//	XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd
//	XTrimMinID(ctx context.Context, key string, minID string) *IntCmd
//	XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd
//	XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd
//	XInfoStream(ctx context.Context, key string) *XInfoStreamCmd
//	XInfoStreamFull(ctx context.Context, key string, count int) *XInfoStreamFullCmd
//	XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd
//
//	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
//	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
//
//	// TODO: remove
//	//		ZAddCh
//	//		ZIncr
//	//		ZAddNXCh
//	//		ZAddXXCh
//	//		ZIncrNX
//	//		ZIncrXX
//	// 	in v9.
//	// 	use ZAddArgs and ZAddArgsIncr.
//
//	ZAdd(ctx context.Context, key string, members ...*Z) *IntCmd
//	ZAddNX(ctx context.Context, key string, members ...*Z) *IntCmd
//	ZAddXX(ctx context.Context, key string, members ...*Z) *IntCmd
//	ZAddCh(ctx context.Context, key string, members ...*Z) *IntCmd
//	ZAddNXCh(ctx context.Context, key string, members ...*Z) *IntCmd
//	ZAddXXCh(ctx context.Context, key string, members ...*Z) *IntCmd
//	ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd
//	ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd
//	ZIncr(ctx context.Context, key string, member *Z) *FloatCmd
//	ZIncrNX(ctx context.Context, key string, member *Z) *FloatCmd
//	ZIncrXX(ctx context.Context, key string, member *Z) *FloatCmd
//	ZCard(ctx context.Context, key string) *IntCmd
//	ZCount(ctx context.Context, key, min, max string) *IntCmd
//	ZLexCount(ctx context.Context, key, min, max string) *IntCmd
//	ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd
//	ZInter(ctx context.Context, store *ZStore) *StringSliceCmd
//	ZInterWithScores(ctx context.Context, store *ZStore) *ZSliceCmd
//	ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd
//	ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd
//	ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd
//	ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd
//	ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
//	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
//	ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
//	ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
//	ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
//	ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd
//	ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd
//	ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd
//	ZRank(ctx context.Context, key, member string) *IntCmd
//	ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd
//	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd
//	ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd
//	ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd
//	ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
//	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
//	ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
//	ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
//	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
//	ZRevRank(ctx context.Context, key, member string) *IntCmd
//	ZScore(ctx context.Context, key, member string) *FloatCmd
//	ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd
//	ZUnion(ctx context.Context, store ZStore) *StringSliceCmd
//	ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd
//	ZRandMember(ctx context.Context, key string, count int, withScores bool) *StringSliceCmd
//	ZDiff(ctx context.Context, keys ...string) *StringSliceCmd
//	ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd
//	ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
//
//	PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd
//	PFCount(ctx context.Context, keys ...string) *IntCmd
//	PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd
//
//	BgRewriteAOF(ctx context.Context) *StatusCmd
//	BgSave(ctx context.Context) *StatusCmd
//	ClientKill(ctx context.Context, ipPort string) *StatusCmd
//	ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd
//	ClientList(ctx context.Context) *StringCmd
//	ClientPause(ctx context.Context, dur time.Duration) *BoolCmd
//	ClientID(ctx context.Context) *IntCmd
//	ConfigGet(ctx context.Context, parameter string) *SliceCmd
//	ConfigResetStat(ctx context.Context) *StatusCmd
//	ConfigSet(ctx context.Context, parameter, value string) *StatusCmd
//	ConfigRewrite(ctx context.Context) *StatusCmd
//	DBSize(ctx context.Context) *IntCmd
//	FlushAll(ctx context.Context) *StatusCmd
//	FlushAllAsync(ctx context.Context) *StatusCmd
//	FlushDB(ctx context.Context) *StatusCmd
//	FlushDBAsync(ctx context.Context) *StatusCmd
//	Info(ctx context.Context, section ...string) *StringCmd
//	LastSave(ctx context.Context) *IntCmd
//	Save(ctx context.Context) *StatusCmd
//	Shutdown(ctx context.Context) *StatusCmd
//	ShutdownSave(ctx context.Context) *StatusCmd
//	ShutdownNoSave(ctx context.Context) *StatusCmd
//	SlaveOf(ctx context.Context, host, port string) *StatusCmd
//	Time(ctx context.Context) *TimeCmd
//	DebugObject(ctx context.Context, key string) *StringCmd
//	ReadOnly(ctx context.Context) *StatusCmd
//	ReadWrite(ctx context.Context) *StatusCmd
//	MemoryUsage(ctx context.Context, key string, samples ...int) *IntCmd
//
//	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
//	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
//	ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd
//	ScriptFlush(ctx context.Context) *StatusCmd
//	ScriptKill(ctx context.Context) *StatusCmd
//	ScriptLoad(ctx context.Context, script string) *StringCmd
//
//	Publish(ctx context.Context, channel string, message interface{}) *IntCmd
//	PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd
//	PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd
//	PubSubNumPat(ctx context.Context) *IntCmd
//
//	ClusterSlots(ctx context.Context) *ClusterSlotsCmd
//	ClusterNodes(ctx context.Context) *StringCmd
//	ClusterMeet(ctx context.Context, host, port string) *StatusCmd
//	ClusterForget(ctx context.Context, nodeID string) *StatusCmd
//	ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd
//	ClusterResetSoft(ctx context.Context) *StatusCmd
//	ClusterResetHard(ctx context.Context) *StatusCmd
//	ClusterInfo(ctx context.Context) *StringCmd
//	ClusterKeySlot(ctx context.Context, key string) *IntCmd
//	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *StringSliceCmd
//	ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd
//	ClusterCountKeysInSlot(ctx context.Context, slot int) *IntCmd
//	ClusterDelSlots(ctx context.Context, slots ...int) *StatusCmd
//	ClusterDelSlotsRange(ctx context.Context, min, max int) *StatusCmd
//	ClusterSaveConfig(ctx context.Context) *StatusCmd
//	ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd
//	ClusterFailover(ctx context.Context) *StatusCmd
//	ClusterAddSlots(ctx context.Context, slots ...int) *StatusCmd
//	ClusterAddSlotsRange(ctx context.Context, min, max int) *StatusCmd
//
//	GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *IntCmd
//	GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
//	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
//	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *IntCmd
//	GeoRadiusByMember(ctx context.Context, key, member string, query *GeoRadiusQuery) *GeoLocationCmd
//	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *GeoRadiusQuery) *IntCmd
//	GeoSearch(ctx context.Context, key string, q *GeoSearchQuery) *StringSliceCmd
//	GeoSearchLocation(ctx context.Context, key string, q *GeoSearchLocationQuery) *GeoSearchLocationCmd
//	GeoSearchStore(ctx context.Context, key, store string, q *GeoSearchStoreQuery) *IntCmd
//	GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd
//	GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd
//}

type StringStringMapCmd = redis.StringStringMapCmd
