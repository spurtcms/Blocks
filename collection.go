package blocks

import (
	"fmt"
	"time"

	"github.com/spurtcms/blocks/migration"
)

// role&permission setup config
func BlockSetup(config Config) *Block {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Block{
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		DB:               config.DB,
		Auth:             config.Auth,
	}

}

/* Collection List*/
// pass limit , offset get collectionlist
func (blocks *Block) CollectionList(filter Filter, tenantid int) (collectionlists []TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	collectionlist, err := Blockmodel.CollectionLists(filter, blocks.DB, tenantid)

	if err != nil {

		fmt.Println(err)
	}

	return collectionlist, nil

}

// Block list
func (blocks *Block) BlockList(filter Filter, tenantid int) (blocklists []TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return []TblBlock{}, AuthErr
	}

	Blockmodel.DataAccess = blocks.DataAccess

	Blockmodel.UserId = blocks.UserId

	blocklist, err := Blockmodel.BlockLists(filter, blocks.DB, tenantid)

	if err != nil {

		fmt.Println(err)
	}

	return blocklist, nil

}

func (blocks *Block) CreateBlock(Bc BlockCreation) (createblocks TblBlock, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlock{}, AuthErr
	}

	var block TblBlock

	block.TenantId = Bc.TenantId
	block.BlockContent = Bc.BlockContent
	block.BlockCss = Bc.BlockCss
	block.BlockDescription = Bc.BlockDescription
	block.CoverImage = Bc.CoverImage
	block.Title = Bc.Title
	block.CreatedBy = Bc.CreatedBy
	block.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	createblock, err := Blockmodel.CreateBlocks(block, blocks.DB)

	if err != nil {

		fmt.Println(err)
	}

	return createblock, nil

}

func (blocks *Block) CheckTagName(tagname string) (flg TblBlockMstrTag, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlockMstrTag{}, AuthErr
	}

	var block TblBlockMstrTag

	tag, err1 := Blockmodel.TagNameCheck(tagname, blocks.DB, block)

	if err1 != nil {
		return TblBlockMstrTag{}, err
	}

	return tag, nil

}

func (blocks *Block) CreateMasterTag(MstrTag MasterTagCreate) (createtags TblBlockMstrTag, err error) {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return TblBlockMstrTag{}, AuthErr
	}

	var tags TblBlockMstrTag

	tags.Name = MstrTag.Name
	tags.TenantId = MstrTag.TenantId
	tags.CreatedBy = MstrTag.CreatedBy
	tags.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	createtag, err := Blockmodel.CreateMasterTag(tags, blocks.DB)

	if err != nil {
		return TblBlockMstrTag{}, err
	}

	return createtag, nil

}

func (blocks *Block) CreateTag(Tag CreateTag) error {

	if AuthErr := AuthandPermission(blocks); AuthErr != nil {

		return AuthErr
	}

	var tags TblBlockTags
	tags.BlockId = Tag.BlockId
	tags.TagId = Tag.TagId
	tags.TagName = Tag.TagName
	tags.TenantId = Tag.TenantId
	tags.CreatedBy = Tag.CreatedBy
	tags.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Blockmodel.CreateBlockTag(tags, blocks.DB)

	if err != nil {
		return err
	}

	return nil

}
