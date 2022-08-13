// Prebuilt snowflake Provider for Terraform CDK (cdktf)
package snowflake

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourceMonitorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifier for the resource monitor; must be unique for your account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#name ResourceMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of credits allocated monthly to the resource monitor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#credit_quota ResourceMonitor#credit_quota}
	CreditQuota *float64 `field:"optional" json:"creditQuota" yaml:"creditQuota"`
	// The date and time when the resource monitor suspends the assigned warehouses.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#end_timestamp ResourceMonitor#end_timestamp}
	EndTimestamp *string `field:"optional" json:"endTimestamp" yaml:"endTimestamp"`
	// The frequency interval at which the credit usage resets to 0.
	//
	// If you set a frequency for a resource monitor, you must also set START_TIMESTAMP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#frequency ResourceMonitor#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#id ResourceMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of percentage thresholds at which to send an alert to subscribed users.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#notify_triggers ResourceMonitor#notify_triggers}
	NotifyTriggers *[]*float64 `field:"optional" json:"notifyTriggers" yaml:"notifyTriggers"`
	// Specifies whether the resource monitor should be applied globally to your Snowflake account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#set_for_account ResourceMonitor#set_for_account}
	SetForAccount interface{} `field:"optional" json:"setForAccount" yaml:"setForAccount"`
	// The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#start_timestamp ResourceMonitor#start_timestamp}
	StartTimestamp *string `field:"optional" json:"startTimestamp" yaml:"startTimestamp"`
	// A list of percentage thresholds at which to immediately suspend all warehouses.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#suspend_immediate_triggers ResourceMonitor#suspend_immediate_triggers}
	SuspendImmediateTriggers *[]*float64 `field:"optional" json:"suspendImmediateTriggers" yaml:"suspendImmediateTriggers"`
	// A list of percentage thresholds at which to suspend all warehouses.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#suspend_triggers ResourceMonitor#suspend_triggers}
	SuspendTriggers *[]*float64 `field:"optional" json:"suspendTriggers" yaml:"suspendTriggers"`
	// A list of warehouses to apply the resource monitor to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/snowflake/r/resource_monitor#warehouses ResourceMonitor#warehouses}
	Warehouses *[]*string `field:"optional" json:"warehouses" yaml:"warehouses"`
}

