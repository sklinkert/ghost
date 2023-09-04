package ghost

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func mustGetCredentialsFromEnv() (ghostURL, ghostContentAPIToken, ghostAdminAPIToken string) {
	ghostURL = os.Getenv("GHOST_URL")
	ghostContentAPIToken = os.Getenv("GHOST_CONTENT_API_TOKEN")
	ghostAdminAPIToken = os.Getenv("GHOST_ADMIN_API_TOKEN")
	if ghostURL == "" || ghostAdminAPIToken == "" || ghostContentAPIToken == "" {
		panic("GHOST_URL, GHOST_ADMIN_API_TOKEN and GHOST_CONTENT_API_TOKEN must be set in the environment")
	}
	return
}

func TestCreateAndDeleteTag(t *testing.T) {
	ghostURL, ghostContentAPIToken, ghostAdminAPIToken := mustGetCredentialsFromEnv()
	g := New(ghostURL, ghostContentAPIToken, ghostAdminAPIToken)

	// Fetch existing tags
	originalTags, err := g.AdminGetTags()
	if err != nil {
		t.Fatalf("Error getting tags: %s", err)
	}

	// Create new tag
	newTag := NewTag{
		Name:            "Test Tag",
		Description:     "This is a test tag",
		Slug:            "test-tag",
		MetaTitle:       "Test Tag",
		MetaDescription: "This is a test tag",
	}
	err = g.AdminCreateTags(NewTags{Tags: []NewTag{newTag}})
	if err != nil {
		t.Fatalf("Error creating tag: %s", err)
	}

	// Fetch tags again
	tagsAfterCreation, err := g.AdminGetTags()
	if err != nil {
		t.Fatalf("Error getting tags second time: %s", err)
	}

	// Check that the new tag is in the list
	found := false
	for _, tag := range tagsAfterCreation.Tags {
		if tag.Name == newTag.Name {
			found = true

			// Delete the tag
			err = g.AdminDeleteTag(tag)
			if err != nil {
				t.Fatalf("Error deleting tag: %s", err)
			}
			break
		}
	}
	if !found {
		t.Fatalf("Tag not found in list after creation")
	}

	tagsAfterDeletion, err := g.AdminGetTags()
	if err != nil {
		t.Fatalf("Error getting tags: %s", err)
	}
	if len(tagsAfterDeletion.Tags) != len(originalTags.Tags) {
		t.Fatalf("Tag count changed after deletion")
	}
}

func TestGetMembers(t *testing.T) {
	ghostURL, ghostContentAPIToken, ghostAdminAPIToken := mustGetCredentialsFromEnv()
	g := New(ghostURL, ghostContentAPIToken, ghostAdminAPIToken)

	members, err := g.AdminGetMembers()
	if err != nil {
		t.Fatalf("Error getting members: %s", err)
	}

	rand.Seed(time.Now().UnixNano())
	randomMailAddress := fmt.Sprintf("testmail-%d@gmx.de", rand.Int())
	_, err = g.AdminCreateMember(NewMember{
		Name:  "Test Member",
		Email: randomMailAddress,
	})
	if err != nil {
		t.Fatalf("Error creating member: %s", err)
	}

	// Fetch members again and compare size of lists
	membersAfterCreation, err := g.AdminGetMembers()
	if err != nil {
		t.Fatalf("Error getting members second time: %s", err)
	}
	if len(membersAfterCreation.Members) == len(members.Members) {
		t.Fatalf("Member count did not change after creation")
	}

	// Check if email matches
	found := false
	for _, member := range membersAfterCreation.Members {
		if member.Email == randomMailAddress {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Member not found in list after creation")
	}
}
