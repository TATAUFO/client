// Auto-generated by avdl-compiler v1.3.20 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/git.avdl

package keybase1

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type EncryptedGitMetadata struct {
	V   int                  `codec:"v" json:"v"`
	E   []byte               `codec:"e" json:"e"`
	N   BoxNonce             `codec:"n" json:"n"`
	Gen PerTeamKeyGeneration `codec:"gen" json:"gen"`
}

func (o EncryptedGitMetadata) DeepCopy() EncryptedGitMetadata {
	return EncryptedGitMetadata{
		V: o.V,
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.E),
		N:   o.N.DeepCopy(),
		Gen: o.Gen.DeepCopy(),
	}
}

type RepoID string

func (o RepoID) DeepCopy() RepoID {
	return o
}

type GitLocalMetadataVersion int

const (
	GitLocalMetadataVersion_V1 GitLocalMetadataVersion = 1
)

func (o GitLocalMetadataVersion) DeepCopy() GitLocalMetadataVersion { return o }

var GitLocalMetadataVersionMap = map[string]GitLocalMetadataVersion{
	"V1": 1,
}

var GitLocalMetadataVersionRevMap = map[GitLocalMetadataVersion]string{
	1: "V1",
}

func (e GitLocalMetadataVersion) String() string {
	if v, ok := GitLocalMetadataVersionRevMap[e]; ok {
		return v
	}
	return ""
}

type GitLocalMetadataV1 struct {
	RepoName GitRepoName `codec:"repoName" json:"repoName"`
}

func (o GitLocalMetadataV1) DeepCopy() GitLocalMetadataV1 {
	return GitLocalMetadataV1{
		RepoName: o.RepoName.DeepCopy(),
	}
}

type GitLocalMetadataVersioned struct {
	Version__ GitLocalMetadataVersion `codec:"version" json:"version"`
	V1__      *GitLocalMetadataV1     `codec:"v1,omitempty" json:"v1,omitempty"`
}

func (o *GitLocalMetadataVersioned) Version() (ret GitLocalMetadataVersion, err error) {
	switch o.Version__ {
	case GitLocalMetadataVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	}
	return o.Version__, nil
}

func (o GitLocalMetadataVersioned) V1() (res GitLocalMetadataV1) {
	if o.Version__ != GitLocalMetadataVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return
	}
	return *o.V1__
}

func NewGitLocalMetadataVersionedWithV1(v GitLocalMetadataV1) GitLocalMetadataVersioned {
	return GitLocalMetadataVersioned{
		Version__: GitLocalMetadataVersion_V1,
		V1__:      &v,
	}
}

func (o GitLocalMetadataVersioned) DeepCopy() GitLocalMetadataVersioned {
	return GitLocalMetadataVersioned{
		Version__: o.Version__.DeepCopy(),
		V1__: (func(x *GitLocalMetadataV1) *GitLocalMetadataV1 {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.V1__),
	}
}

type GitLocalMetadata struct {
	RepoName GitRepoName `codec:"repoName" json:"repoName"`
}

func (o GitLocalMetadata) DeepCopy() GitLocalMetadata {
	return GitLocalMetadata{
		RepoName: o.RepoName.DeepCopy(),
	}
}

type GitServerMetadata struct {
	Ctime                   Time     `codec:"ctime" json:"ctime"`
	Mtime                   Time     `codec:"mtime" json:"mtime"`
	LastModifyingUsername   string   `codec:"lastModifyingUsername" json:"lastModifyingUsername"`
	LastModifyingDeviceID   DeviceID `codec:"lastModifyingDeviceID" json:"lastModifyingDeviceID"`
	LastModifyingDeviceName string   `codec:"lastModifyingDeviceName" json:"lastModifyingDeviceName"`
}

func (o GitServerMetadata) DeepCopy() GitServerMetadata {
	return GitServerMetadata{
		Ctime: o.Ctime.DeepCopy(),
		Mtime: o.Mtime.DeepCopy(),
		LastModifyingUsername:   o.LastModifyingUsername,
		LastModifyingDeviceID:   o.LastModifyingDeviceID.DeepCopy(),
		LastModifyingDeviceName: o.LastModifyingDeviceName,
	}
}

type GitRepoResult struct {
	Folder         Folder            `codec:"folder" json:"folder"`
	RepoID         RepoID            `codec:"repoID" json:"repoID"`
	LocalMetadata  GitLocalMetadata  `codec:"localMetadata" json:"localMetadata"`
	ServerMetadata GitServerMetadata `codec:"serverMetadata" json:"serverMetadata"`
	RepoUrl        string            `codec:"repoUrl" json:"repoUrl"`
	GlobalUniqueID string            `codec:"globalUniqueID" json:"globalUniqueID"`
	CanDelete      bool              `codec:"canDelete" json:"canDelete"`
}

func (o GitRepoResult) DeepCopy() GitRepoResult {
	return GitRepoResult{
		Folder:         o.Folder.DeepCopy(),
		RepoID:         o.RepoID.DeepCopy(),
		LocalMetadata:  o.LocalMetadata.DeepCopy(),
		ServerMetadata: o.ServerMetadata.DeepCopy(),
		RepoUrl:        o.RepoUrl,
		GlobalUniqueID: o.GlobalUniqueID,
		CanDelete:      o.CanDelete,
	}
}

type PutGitMetadataArg struct {
	Folder     Folder           `codec:"folder" json:"folder"`
	RepoID     RepoID           `codec:"repoID" json:"repoID"`
	Metadata   GitLocalMetadata `codec:"metadata" json:"metadata"`
	NotifyTeam bool             `codec:"notifyTeam" json:"notifyTeam"`
}

func (o PutGitMetadataArg) DeepCopy() PutGitMetadataArg {
	return PutGitMetadataArg{
		Folder:     o.Folder.DeepCopy(),
		RepoID:     o.RepoID.DeepCopy(),
		Metadata:   o.Metadata.DeepCopy(),
		NotifyTeam: o.NotifyTeam,
	}
}

type DeleteGitMetadataArg struct {
	Folder   Folder      `codec:"folder" json:"folder"`
	RepoName GitRepoName `codec:"repoName" json:"repoName"`
}

func (o DeleteGitMetadataArg) DeepCopy() DeleteGitMetadataArg {
	return DeleteGitMetadataArg{
		Folder:   o.Folder.DeepCopy(),
		RepoName: o.RepoName.DeepCopy(),
	}
}

type GetGitMetadataArg struct {
	Folder Folder `codec:"folder" json:"folder"`
}

func (o GetGitMetadataArg) DeepCopy() GetGitMetadataArg {
	return GetGitMetadataArg{
		Folder: o.Folder.DeepCopy(),
	}
}

type GetAllGitMetadataArg struct {
}

func (o GetAllGitMetadataArg) DeepCopy() GetAllGitMetadataArg {
	return GetAllGitMetadataArg{}
}

type CreatePersonalRepoArg struct {
	RepoName GitRepoName `codec:"repoName" json:"repoName"`
}

func (o CreatePersonalRepoArg) DeepCopy() CreatePersonalRepoArg {
	return CreatePersonalRepoArg{
		RepoName: o.RepoName.DeepCopy(),
	}
}

type CreateTeamRepoArg struct {
	RepoName   GitRepoName `codec:"repoName" json:"repoName"`
	TeamName   TeamName    `codec:"teamName" json:"teamName"`
	NotifyTeam bool        `codec:"notifyTeam" json:"notifyTeam"`
}

func (o CreateTeamRepoArg) DeepCopy() CreateTeamRepoArg {
	return CreateTeamRepoArg{
		RepoName:   o.RepoName.DeepCopy(),
		TeamName:   o.TeamName.DeepCopy(),
		NotifyTeam: o.NotifyTeam,
	}
}

type DeletePersonalRepoArg struct {
	RepoName GitRepoName `codec:"repoName" json:"repoName"`
}

func (o DeletePersonalRepoArg) DeepCopy() DeletePersonalRepoArg {
	return DeletePersonalRepoArg{
		RepoName: o.RepoName.DeepCopy(),
	}
}

type DeleteTeamRepoArg struct {
	RepoName   GitRepoName `codec:"repoName" json:"repoName"`
	TeamName   TeamName    `codec:"teamName" json:"teamName"`
	NotifyTeam bool        `codec:"notifyTeam" json:"notifyTeam"`
}

func (o DeleteTeamRepoArg) DeepCopy() DeleteTeamRepoArg {
	return DeleteTeamRepoArg{
		RepoName:   o.RepoName.DeepCopy(),
		TeamName:   o.TeamName.DeepCopy(),
		NotifyTeam: o.NotifyTeam,
	}
}

type GitInterface interface {
	PutGitMetadata(context.Context, PutGitMetadataArg) error
	DeleteGitMetadata(context.Context, DeleteGitMetadataArg) error
	GetGitMetadata(context.Context, Folder) ([]GitRepoResult, error)
	GetAllGitMetadata(context.Context) ([]GitRepoResult, error)
	CreatePersonalRepo(context.Context, GitRepoName) (RepoID, error)
	CreateTeamRepo(context.Context, CreateTeamRepoArg) (RepoID, error)
	DeletePersonalRepo(context.Context, GitRepoName) error
	DeleteTeamRepo(context.Context, DeleteTeamRepoArg) error
}

func GitProtocol(i GitInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.git",
		Methods: map[string]rpc.ServeHandlerDescription{
			"putGitMetadata": {
				MakeArg: func() interface{} {
					ret := make([]PutGitMetadataArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PutGitMetadataArg)
					if !ok {
						err = rpc.NewTypeError((*[]PutGitMetadataArg)(nil), args)
						return
					}
					err = i.PutGitMetadata(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"deleteGitMetadata": {
				MakeArg: func() interface{} {
					ret := make([]DeleteGitMetadataArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DeleteGitMetadataArg)
					if !ok {
						err = rpc.NewTypeError((*[]DeleteGitMetadataArg)(nil), args)
						return
					}
					err = i.DeleteGitMetadata(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getGitMetadata": {
				MakeArg: func() interface{} {
					ret := make([]GetGitMetadataArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetGitMetadataArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetGitMetadataArg)(nil), args)
						return
					}
					ret, err = i.GetGitMetadata(ctx, (*typedArgs)[0].Folder)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getAllGitMetadata": {
				MakeArg: func() interface{} {
					ret := make([]GetAllGitMetadataArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetAllGitMetadata(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"createPersonalRepo": {
				MakeArg: func() interface{} {
					ret := make([]CreatePersonalRepoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CreatePersonalRepoArg)
					if !ok {
						err = rpc.NewTypeError((*[]CreatePersonalRepoArg)(nil), args)
						return
					}
					ret, err = i.CreatePersonalRepo(ctx, (*typedArgs)[0].RepoName)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"createTeamRepo": {
				MakeArg: func() interface{} {
					ret := make([]CreateTeamRepoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CreateTeamRepoArg)
					if !ok {
						err = rpc.NewTypeError((*[]CreateTeamRepoArg)(nil), args)
						return
					}
					ret, err = i.CreateTeamRepo(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"deletePersonalRepo": {
				MakeArg: func() interface{} {
					ret := make([]DeletePersonalRepoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DeletePersonalRepoArg)
					if !ok {
						err = rpc.NewTypeError((*[]DeletePersonalRepoArg)(nil), args)
						return
					}
					err = i.DeletePersonalRepo(ctx, (*typedArgs)[0].RepoName)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"deleteTeamRepo": {
				MakeArg: func() interface{} {
					ret := make([]DeleteTeamRepoArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DeleteTeamRepoArg)
					if !ok {
						err = rpc.NewTypeError((*[]DeleteTeamRepoArg)(nil), args)
						return
					}
					err = i.DeleteTeamRepo(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type GitClient struct {
	Cli rpc.GenericClient
}

func (c GitClient) PutGitMetadata(ctx context.Context, __arg PutGitMetadataArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.git.putGitMetadata", []interface{}{__arg}, nil)
	return
}

func (c GitClient) DeleteGitMetadata(ctx context.Context, __arg DeleteGitMetadataArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.git.deleteGitMetadata", []interface{}{__arg}, nil)
	return
}

func (c GitClient) GetGitMetadata(ctx context.Context, folder Folder) (res []GitRepoResult, err error) {
	__arg := GetGitMetadataArg{Folder: folder}
	err = c.Cli.Call(ctx, "keybase.1.git.getGitMetadata", []interface{}{__arg}, &res)
	return
}

func (c GitClient) GetAllGitMetadata(ctx context.Context) (res []GitRepoResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.git.getAllGitMetadata", []interface{}{GetAllGitMetadataArg{}}, &res)
	return
}

func (c GitClient) CreatePersonalRepo(ctx context.Context, repoName GitRepoName) (res RepoID, err error) {
	__arg := CreatePersonalRepoArg{RepoName: repoName}
	err = c.Cli.Call(ctx, "keybase.1.git.createPersonalRepo", []interface{}{__arg}, &res)
	return
}

func (c GitClient) CreateTeamRepo(ctx context.Context, __arg CreateTeamRepoArg) (res RepoID, err error) {
	err = c.Cli.Call(ctx, "keybase.1.git.createTeamRepo", []interface{}{__arg}, &res)
	return
}

func (c GitClient) DeletePersonalRepo(ctx context.Context, repoName GitRepoName) (err error) {
	__arg := DeletePersonalRepoArg{RepoName: repoName}
	err = c.Cli.Call(ctx, "keybase.1.git.deletePersonalRepo", []interface{}{__arg}, nil)
	return
}

func (c GitClient) DeleteTeamRepo(ctx context.Context, __arg DeleteTeamRepoArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.git.deleteTeamRepo", []interface{}{__arg}, nil)
	return
}
