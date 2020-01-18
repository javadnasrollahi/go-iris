//
// ──────────────────────────────────────────────────────────────────────────────  ──────────
//   :::::: D O N T   E D I T   T H I S   F I L E : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────────────────────
//

/*eslint-disable */
package acl

type AclVal struct {
	Value      uint64 `json:"value"`
	Title      string `json:"title"`
	Comment    string `json:"comment"`
	Pridicates []string
	Critically float32 `json:"critically,omitempty"`
}

type ACL struct {
	ACLDef map[string]AclVal
}

//
// ────────────────────────────────────────────────────────────────────────────── I ──────────
//   :::::: C O N S T A N T   A C L   V A L U E S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────────────────────
//
const GLike = 0x1
const GComment = 0x2
const GPromote = 0x4
const GSendMSG = 0x8
const GRemoveMSG = 0x10
const GDemote = 0x20
const GAddUser = 0x40
const GRemoveUser = 0x80
const GEditChatInfo = 0x100

var GroupAcl = map[string]AclVal{
	"Like": AclVal{
		Value:   GLike,
		Title:   "Like",
		Comment: "Like",
	},
	"Comment": AclVal{
		Value:   GComment,
		Title:   "Comment",
		Comment: "Comment",
	},
	"Promote": AclVal{
		Value:   GPromote,
		Title:   "Promote",
		Comment: "Promote",
	},
	"SendMSG": AclVal{
		Value:   GSendMSG,
		Title:   "SendMSG",
		Comment: "SendMSG",
	},
	"RemoveMSG": AclVal{
		Value:   GRemoveMSG,
		Title:   "RemoveMSG",
		Comment: "RemoveMSG",
	},
	"Demote": AclVal{
		Value:   GDemote,
		Title:   "Demote",
		Comment: "Demote",
	},
	"AddUser": AclVal{
		Value:   GAddUser,
		Title:   "AddUser",
		Comment: "AddUser",
	},
	"RemoveUser": AclVal{
		Value:   GRemoveUser,
		Title:   "RemoveUser",
		Comment: "RemoveUser",
	},
	"EditChatInfo": AclVal{
		Value:   GEditChatInfo,
		Title:   "EditChatInfo",
		Comment: "EditChatInfo",
	},
}

const View = 0x1
const Comment = 0x2
const Edit = 0x4
const delete = 0x8
const ExamsLocal = 0x20
const ExamsPublic = 0x40
const ExamsSuperView = 0x80
const BookCreate = 0x400
const BookEdit = 0x800
const BookSuperEdit = 0x1000
const BookDelete = 0x2000
const BookSuperDelete = 0x4000
const Admin = 0x20000
const PersonsSuperView = 0x100000
const PersonsSuperEdit = 0x200000

//
var ACLList = map[string]AclVal{

	"View": AclVal{
		Value:   0x1,
		Title:   "مشاهده",
		Comment: "مشاهده",
	},

	"Comment": AclVal{
		Value:   0x2,
		Title:   "کامنت گذاشتن",
		Comment: "کامنت گذاشتن",
	},

	"Edit": AclVal{
		Value:   0x4,
		Title:   "ویرایش",
		Comment: "ویرایش",
	},

	"delete": AclVal{
		Value:   0x8,
		Title:   "حذف",
		Comment: "حذف",
	},
}

var UserDataAL = map[string]AclVal{
	"basic": AclVal{
		Title:   "اطلاعات پایه",
		Comment: "ایمیل ، نام جایگزین",
		Value:   0x1,
		Pridicates: []string{
			"alternateName",
			"email",
		},
		Critically: 0.1,
	},
	"public": AclVal{
		Title:   "اطلاعات هویت عمومی",
		Comment: "موبایل- نام - نام خانوادگی - جنسیت",
		Value:   0x2,
		Pridicates: []string{
			"gender",
			"familyName",
			"mobile",
			"name",
		},
		Critically: 0.2,
	},
	"publicExtra": AclVal{
		Title:   "اطلاعات هویت عمومی",
		Comment: "موبایل- نام - نام خانوادگی - بیوگرافی",
		Value:   0x8,
		Pridicates: []string{
			"bio",
		},
		Critically: 0.2,
	},
}

//SuperAdminACL := 0xffffffffffff
//DefaultUser := 0xf00f1d7f
func Allow(userAcl uint64, acl AclVal) bool {
	return userAcl&acl.Value != 0
}

func AllowCode(userAcl uint64, acl uint64) bool {
	return userAcl&acl != 0
}

func Privileges(Useracl uint64) map[string]AclVal {
	//ctx.Writef("%#x", Useracl)
	res := make(map[string]AclVal)
	for pname, pel := range ACLList {
		if Allow(Useracl, pel) {
			res[pname] = AclVal{}
			res[pname] = pel
		}
	}
	return res
}
func UserDataPrivileges(Useracl uint64) map[string]AclVal {
	res := make(map[string]AclVal)
	for pname, pel := range UserDataAL {
		if Allow(Useracl, pel) {
			res[pname] = AclVal{}
			res[pname] = pel

		}
	}
	return res
}
