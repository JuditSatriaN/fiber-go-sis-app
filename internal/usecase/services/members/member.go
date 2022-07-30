package members

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	memberEntity "github.com/fiber-go-sis-app/internal/entity/members"

	personaliaRepo "github.com/fiber-go-sis-app/internal/repo/members"
)

func GetAllMember(ctx *fiber.Ctx) ([]memberEntity.Member, error) {
	members, err := personaliaRepo.GetAllMemberRepo(ctx)
	if err != nil {
		return []memberEntity.Member{}, err
	}

	return members, nil
}

func GetMemberByID(ctx *fiber.Ctx, ID int) (memberEntity.Member, error) {
	member, found, err := personaliaRepo.GetMemberByID(ctx, ID)
	if err != nil {
		return memberEntity.Member{}, err
	}

	if !found {
		return memberEntity.Member{}, fmt.Errorf("member dengan nama : %s tidak ditemukan", member.Name)
	}

	return member, nil
}

func InsertMember(ctx *fiber.Ctx, member memberEntity.Member) error {
	return personaliaRepo.InsertMember(ctx, member)
}

func UpdateMember(ctx *fiber.Ctx, member memberEntity.Member) error {
	if _, err := GetMemberByID(ctx, member.ID); err != nil {
		return err
	}

	return personaliaRepo.UpdateMember(ctx, member)
}

func DeleteMember(ctx *fiber.Ctx, ID int) error {
	if _, err := GetMemberByID(ctx, ID); err != nil {
		return err
	}

	return personaliaRepo.DeleteMember(ctx, ID)
}

func UpsertMember(ctx *fiber.Ctx, member memberEntity.Member) error {
	_, found, err := personaliaRepo.GetMemberByID(ctx, member.ID)
	if err != nil {
		return err
	}

	if !found || member.ID == 0 {
		return personaliaRepo.InsertMember(ctx, member)
	} else {
		return personaliaRepo.UpdateMember(ctx, member)
	}
}
