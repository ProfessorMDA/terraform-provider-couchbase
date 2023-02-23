package couchbase

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBuildIndex() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBuildIndexCreate,
		DeleteContext: resourceBuildIndexDelete,

		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			//"index_name": {
			//	Type:     schema.TypeString,
			//	Required: true,
			//},
		},
	}
}

func resourceBuildIndexCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	couchbase, diags := m.(*Connection).CouchbaseInitialization()
	if diags != nil {
		return diags
	}
	defer couchbase.ConnectionCLose()

	bucketName := d.Get("bucket_name").(string)
	//indexName := d.Get("index_name").(string)

	fmt.Sprintf("BUILD INDEX ON '%s'(( SELECT RAW name FROM system:indexes where keyspace_id = '%s' AND state = 'deferred' )) USING GSI", bucketName, bucketName)
	//_, err := cc.Cluster.Query(q, nil)
	//if err != nil {
	//	return err
	//}

	// Set the ID of the created resource to the name of the bucket
	d.SetId(bucketName)
	return nil
}

func resourceBuildIndexDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Deleting a resource doesn't make sense in this case, since building indexes is idempotent.
	return nil
}
