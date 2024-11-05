package host

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/sendsmaily/terraform-provider-definednet/internal/definednet"
)

// State is the host resource's state.
type State struct {
	ID             types.String `tfsdk:"id"`
	NetworkID      types.String `tfsdk:"network_id"`
	RoleID         types.String `tfsdk:"role_id"`
	Name           types.String `tfsdk:"name"`
	IPAddress      types.String `tfsdk:"ip_address"`
	Tags           types.List   `tfsdk:"tags"`
	EnrollmentCode types.String `tfsdk:"enrollment_code"`
	Metrics        types.Object `tfsdk:"metrics"`
}

// ApplyEnrollment applies Defined.net host enrollment information to the state.
func (s *State) ApplyEnrollment(ctx context.Context, enrollment *definednet.Enrollment) (diags diag.Diagnostics) {
	diags.Append(s.ApplyHost(ctx, &enrollment.Host)...)
	s.EnrollmentCode = types.StringValue(enrollment.EnrollmentCode.Code)

	return diags
}

// ApplyHost applies Defined.net host information to the state.
func (s *State) ApplyHost(ctx context.Context, host *definednet.Host) (diags diag.Diagnostics) {
	s.ID = types.StringValue(host.ID)
	s.IPAddress = types.StringValue(host.IPAddress)
	s.Name = types.StringValue(host.Name)
	s.NetworkID = types.StringValue(host.NetworkID)
	s.RoleID = types.StringValue(host.RoleID)

	{
		var d diag.Diagnostics
		s.Tags, d = types.ListValueFrom(ctx, types.StringType, host.Tags)
		diags = append(diags, d...)
	}

	{
		var d diag.Diagnostics
		s.Metrics, d = types.ObjectValueFrom(ctx, map[string]attr.Type{
			"listen":               basetypes.StringType{},
			"path":                 basetypes.StringType{},
			"namespace":            basetypes.StringType{},
			"subsystem":            basetypes.StringType{},
			"enable_extra_metrics": basetypes.BoolType{},
		}, host.Metrics)
		diags = append(diags, d...)
	}

	return diags
}
