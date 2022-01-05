package handler

import (
	"context"
	"fmt"
	"io"

	pb "github.com/2637309949/micro/v3/proto/runtime"
	"github.com/2637309949/micro/v3/service/auth"
	"github.com/2637309949/micro/v3/service/errors"
	"github.com/2637309949/micro/v3/service/store"
)

const bufferSize = 1024

// Build implements the proto build service interface
type Build struct{}

func (b *Build) Read(ctx context.Context, req *pb.Service, stream pb.Build_ReadStream) error {
	defer stream.Close()

	// authorize the request
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("An account is required to read builds")
	}

	// validate the request
	if len(req.Name) == 0 {
		return errors.BadRequest("Missing name")
	}
	if len(req.Version) == 0 {
		return errors.BadRequest("Missing version")
	}

	// lookup the build from the blob store
	key := fmt.Sprintf("build://%v:%v", req.Name, req.Version)
	build, err := store.DefaultBlobStore.Read(key, store.BlobNamespace(acc.Issuer))
	if err == store.ErrNotFound {
		return errors.NotFound("Build not found")
	} else if err != nil {
		return err
	}

	// read bytes from the store and stream it to the client
	buffer := make([]byte, bufferSize)
	for {
		num, err := build.Read(buffer)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return errors.InternalServerError("Error reading build from store: %v", err)
		}

		if err := stream.Send(&pb.BuildReadResponse{Data: buffer[:num]}); err != nil {
			return err
		}
	}
}
