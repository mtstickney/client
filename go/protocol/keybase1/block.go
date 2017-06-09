// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/block.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type GetBlockRes struct {
	BlockKey string `codec:"blockKey" json:"blockKey"`
	Buf      []byte `codec:"buf" json:"buf"`
}

func (o GetBlockRes) DeepCopy() GetBlockRes {
	return GetBlockRes{
		BlockKey: o.BlockKey,
		Buf: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.Buf),
	}
}

type BlockRefNonce [8]byte

func (o BlockRefNonce) DeepCopy() BlockRefNonce {
	var ret BlockRefNonce
	copy(ret[:], o[:])
	return ret
}

type BlockReference struct {
	Bid       BlockIdCombo  `codec:"bid" json:"bid"`
	Nonce     BlockRefNonce `codec:"nonce" json:"nonce"`
	ChargedTo UserOrTeamID  `codec:"chargedTo" json:"chargedTo"`
}

func (o BlockReference) DeepCopy() BlockReference {
	return BlockReference{
		Bid:       o.Bid.DeepCopy(),
		Nonce:     o.Nonce.DeepCopy(),
		ChargedTo: o.ChargedTo.DeepCopy(),
	}
}

type BlockReferenceCount struct {
	Ref       BlockReference `codec:"ref" json:"ref"`
	LiveCount int            `codec:"liveCount" json:"liveCount"`
}

func (o BlockReferenceCount) DeepCopy() BlockReferenceCount {
	return BlockReferenceCount{
		Ref:       o.Ref.DeepCopy(),
		LiveCount: o.LiveCount,
	}
}

type DowngradeReferenceRes struct {
	Completed []BlockReferenceCount `codec:"completed" json:"completed"`
	Failed    BlockReference        `codec:"failed" json:"failed"`
}

func (o DowngradeReferenceRes) DeepCopy() DowngradeReferenceRes {
	return DowngradeReferenceRes{
		Completed: (func(x []BlockReferenceCount) []BlockReferenceCount {
			var ret []BlockReferenceCount
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Completed),
		Failed: o.Failed.DeepCopy(),
	}
}

type BlockPingResponse struct {
}

func (o BlockPingResponse) DeepCopy() BlockPingResponse {
	return BlockPingResponse{}
}

type GetSessionChallengeArg struct {
}

func (o GetSessionChallengeArg) DeepCopy() GetSessionChallengeArg {
	return GetSessionChallengeArg{}
}

type AuthenticateSessionArg struct {
	Signature string `codec:"signature" json:"signature"`
}

func (o AuthenticateSessionArg) DeepCopy() AuthenticateSessionArg {
	return AuthenticateSessionArg{
		Signature: o.Signature,
	}
}

type PutBlockArg struct {
	Bid      BlockIdCombo `codec:"bid" json:"bid"`
	Folder   string       `codec:"folder" json:"folder"`
	BlockKey string       `codec:"blockKey" json:"blockKey"`
	Buf      []byte       `codec:"buf" json:"buf"`
}

func (o PutBlockArg) DeepCopy() PutBlockArg {
	return PutBlockArg{
		Bid:      o.Bid.DeepCopy(),
		Folder:   o.Folder,
		BlockKey: o.BlockKey,
		Buf: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.Buf),
	}
}

type PutBlockAgainArg struct {
	Folder   string         `codec:"folder" json:"folder"`
	Ref      BlockReference `codec:"ref" json:"ref"`
	BlockKey string         `codec:"blockKey" json:"blockKey"`
	Buf      []byte         `codec:"buf" json:"buf"`
}

func (o PutBlockAgainArg) DeepCopy() PutBlockAgainArg {
	return PutBlockAgainArg{
		Folder:   o.Folder,
		Ref:      o.Ref.DeepCopy(),
		BlockKey: o.BlockKey,
		Buf: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte(nil), x...)
		})(o.Buf),
	}
}

type GetBlockArg struct {
	Bid    BlockIdCombo `codec:"bid" json:"bid"`
	Folder string       `codec:"folder" json:"folder"`
}

func (o GetBlockArg) DeepCopy() GetBlockArg {
	return GetBlockArg{
		Bid:    o.Bid.DeepCopy(),
		Folder: o.Folder,
	}
}

type AddReferenceArg struct {
	Folder string         `codec:"folder" json:"folder"`
	Ref    BlockReference `codec:"ref" json:"ref"`
}

func (o AddReferenceArg) DeepCopy() AddReferenceArg {
	return AddReferenceArg{
		Folder: o.Folder,
		Ref:    o.Ref.DeepCopy(),
	}
}

type DelReferenceArg struct {
	Folder string         `codec:"folder" json:"folder"`
	Ref    BlockReference `codec:"ref" json:"ref"`
}

func (o DelReferenceArg) DeepCopy() DelReferenceArg {
	return DelReferenceArg{
		Folder: o.Folder,
		Ref:    o.Ref.DeepCopy(),
	}
}

type ArchiveReferenceArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

func (o ArchiveReferenceArg) DeepCopy() ArchiveReferenceArg {
	return ArchiveReferenceArg{
		Folder: o.Folder,
		Refs: (func(x []BlockReference) []BlockReference {
			var ret []BlockReference
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Refs),
	}
}

type DelReferenceWithCountArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

func (o DelReferenceWithCountArg) DeepCopy() DelReferenceWithCountArg {
	return DelReferenceWithCountArg{
		Folder: o.Folder,
		Refs: (func(x []BlockReference) []BlockReference {
			var ret []BlockReference
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Refs),
	}
}

type ArchiveReferenceWithCountArg struct {
	Folder string           `codec:"folder" json:"folder"`
	Refs   []BlockReference `codec:"refs" json:"refs"`
}

func (o ArchiveReferenceWithCountArg) DeepCopy() ArchiveReferenceWithCountArg {
	return ArchiveReferenceWithCountArg{
		Folder: o.Folder,
		Refs: (func(x []BlockReference) []BlockReference {
			var ret []BlockReference
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Refs),
	}
}

type GetUserQuotaInfoArg struct {
}

func (o GetUserQuotaInfoArg) DeepCopy() GetUserQuotaInfoArg {
	return GetUserQuotaInfoArg{}
}

type GetTeamQuotaInfoArg struct {
	Tid TeamID `codec:"tid" json:"tid"`
}

func (o GetTeamQuotaInfoArg) DeepCopy() GetTeamQuotaInfoArg {
	return GetTeamQuotaInfoArg{
		Tid: o.Tid.DeepCopy(),
	}
}

type BlockPingArg struct {
}

func (o BlockPingArg) DeepCopy() BlockPingArg {
	return BlockPingArg{}
}

type BlockInterface interface {
	GetSessionChallenge(context.Context) (ChallengeInfo, error)
	AuthenticateSession(context.Context, string) error
	PutBlock(context.Context, PutBlockArg) error
	PutBlockAgain(context.Context, PutBlockAgainArg) error
	GetBlock(context.Context, GetBlockArg) (GetBlockRes, error)
	AddReference(context.Context, AddReferenceArg) error
	DelReference(context.Context, DelReferenceArg) error
	ArchiveReference(context.Context, ArchiveReferenceArg) ([]BlockReference, error)
	DelReferenceWithCount(context.Context, DelReferenceWithCountArg) (DowngradeReferenceRes, error)
	ArchiveReferenceWithCount(context.Context, ArchiveReferenceWithCountArg) (DowngradeReferenceRes, error)
	GetUserQuotaInfo(context.Context) ([]byte, error)
	GetTeamQuotaInfo(context.Context, TeamID) ([]byte, error)
	BlockPing(context.Context) (BlockPingResponse, error)
}

func BlockProtocol(i BlockInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.block",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getSessionChallenge": {
				MakeArg: func() interface{} {
					ret := make([]GetSessionChallengeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetSessionChallenge(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"authenticateSession": {
				MakeArg: func() interface{} {
					ret := make([]AuthenticateSessionArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AuthenticateSessionArg)
					if !ok {
						err = rpc.NewTypeError((*[]AuthenticateSessionArg)(nil), args)
						return
					}
					err = i.AuthenticateSession(ctx, (*typedArgs)[0].Signature)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"putBlock": {
				MakeArg: func() interface{} {
					ret := make([]PutBlockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PutBlockArg)
					if !ok {
						err = rpc.NewTypeError((*[]PutBlockArg)(nil), args)
						return
					}
					err = i.PutBlock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"putBlockAgain": {
				MakeArg: func() interface{} {
					ret := make([]PutBlockAgainArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PutBlockAgainArg)
					if !ok {
						err = rpc.NewTypeError((*[]PutBlockAgainArg)(nil), args)
						return
					}
					err = i.PutBlockAgain(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getBlock": {
				MakeArg: func() interface{} {
					ret := make([]GetBlockArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetBlockArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetBlockArg)(nil), args)
						return
					}
					ret, err = i.GetBlock(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"addReference": {
				MakeArg: func() interface{} {
					ret := make([]AddReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]AddReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]AddReferenceArg)(nil), args)
						return
					}
					err = i.AddReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"delReference": {
				MakeArg: func() interface{} {
					ret := make([]DelReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DelReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]DelReferenceArg)(nil), args)
						return
					}
					err = i.DelReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"archiveReference": {
				MakeArg: func() interface{} {
					ret := make([]ArchiveReferenceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ArchiveReferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[]ArchiveReferenceArg)(nil), args)
						return
					}
					ret, err = i.ArchiveReference(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"delReferenceWithCount": {
				MakeArg: func() interface{} {
					ret := make([]DelReferenceWithCountArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DelReferenceWithCountArg)
					if !ok {
						err = rpc.NewTypeError((*[]DelReferenceWithCountArg)(nil), args)
						return
					}
					ret, err = i.DelReferenceWithCount(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"archiveReferenceWithCount": {
				MakeArg: func() interface{} {
					ret := make([]ArchiveReferenceWithCountArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ArchiveReferenceWithCountArg)
					if !ok {
						err = rpc.NewTypeError((*[]ArchiveReferenceWithCountArg)(nil), args)
						return
					}
					ret, err = i.ArchiveReferenceWithCount(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getUserQuotaInfo": {
				MakeArg: func() interface{} {
					ret := make([]GetUserQuotaInfoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetUserQuotaInfo(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getTeamQuotaInfo": {
				MakeArg: func() interface{} {
					ret := make([]GetTeamQuotaInfoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetTeamQuotaInfoArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetTeamQuotaInfoArg)(nil), args)
						return
					}
					ret, err = i.GetTeamQuotaInfo(ctx, (*typedArgs)[0].Tid)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"blockPing": {
				MakeArg: func() interface{} {
					ret := make([]BlockPingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.BlockPing(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type BlockClient struct {
	Cli rpc.GenericClient
}

func (c BlockClient) GetSessionChallenge(ctx context.Context) (res ChallengeInfo, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getSessionChallenge", []interface{}{GetSessionChallengeArg{}}, &res)
	return
}

func (c BlockClient) AuthenticateSession(ctx context.Context, signature string) (err error) {
	__arg := AuthenticateSessionArg{Signature: signature}
	err = c.Cli.Call(ctx, "keybase.1.block.authenticateSession", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) PutBlock(ctx context.Context, __arg PutBlockArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.putBlock", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) PutBlockAgain(ctx context.Context, __arg PutBlockAgainArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.putBlockAgain", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) GetBlock(ctx context.Context, __arg GetBlockArg) (res GetBlockRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getBlock", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) AddReference(ctx context.Context, __arg AddReferenceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.addReference", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) DelReference(ctx context.Context, __arg DelReferenceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.delReference", []interface{}{__arg}, nil)
	return
}

func (c BlockClient) ArchiveReference(ctx context.Context, __arg ArchiveReferenceArg) (res []BlockReference, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.archiveReference", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) DelReferenceWithCount(ctx context.Context, __arg DelReferenceWithCountArg) (res DowngradeReferenceRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.delReferenceWithCount", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) ArchiveReferenceWithCount(ctx context.Context, __arg ArchiveReferenceWithCountArg) (res DowngradeReferenceRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.archiveReferenceWithCount", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) GetUserQuotaInfo(ctx context.Context) (res []byte, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.getUserQuotaInfo", []interface{}{GetUserQuotaInfoArg{}}, &res)
	return
}

func (c BlockClient) GetTeamQuotaInfo(ctx context.Context, tid TeamID) (res []byte, err error) {
	__arg := GetTeamQuotaInfoArg{Tid: tid}
	err = c.Cli.Call(ctx, "keybase.1.block.getTeamQuotaInfo", []interface{}{__arg}, &res)
	return
}

func (c BlockClient) BlockPing(ctx context.Context) (res BlockPingResponse, err error) {
	err = c.Cli.Call(ctx, "keybase.1.block.blockPing", []interface{}{BlockPingArg{}}, &res)
	return
}
