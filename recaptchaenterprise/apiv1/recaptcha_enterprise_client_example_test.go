// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recaptchaenterprise_test

import (
	"context"

	recaptchaenterprise "cloud.google.com/go/recaptchaenterprise/apiv1"
	"google.golang.org/api/iterator"
	recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateAssessment() {
	// import recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"

	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recaptchaenterprisepb.CreateAssessmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateAssessment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_AnnotateAssessment() {
	// import recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"

	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recaptchaenterprisepb.AnnotateAssessmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AnnotateAssessment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateKey() {
	// import recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"

	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recaptchaenterprisepb.CreateKeyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListKeys() {
	// import recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recaptchaenterprisepb.ListKeysRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListKeys(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetKey() {
	// import recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"

	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recaptchaenterprisepb.GetKeyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateKey() {
	// import recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1"

	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recaptchaenterprisepb.UpdateKeyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteKey() {
	ctx := context.Background()
	c, err := recaptchaenterprise.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recaptchaenterprisepb.DeleteKeyRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteKey(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
