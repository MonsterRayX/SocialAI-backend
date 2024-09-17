// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/actions/sdk/v2/settings.proto

package sdk

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The category choices for an Actions project.
type Settings_Category int32

const (
	// Unknown / Unspecified.
	Settings_CATEGORY_UNSPECIFIED Settings_Category = 0
	// Business and Finance category.
	Settings_BUSINESS_AND_FINANCE Settings_Category = 2
	// Education and Reference category.
	Settings_EDUCATION_AND_REFERENCE Settings_Category = 3
	// Food and Drink category.
	Settings_FOOD_AND_DRINK Settings_Category = 4
	// Games and Trivia category.
	Settings_GAMES_AND_TRIVIA Settings_Category = 5
	// Health and Fitness category.
	Settings_HEALTH_AND_FITNESS Settings_Category = 6
	// Kids and Family category.
	Settings_KIDS_AND_FAMILY Settings_Category = 20
	// Lifestyle category.
	Settings_LIFESTYLE Settings_Category = 7
	// Local category.
	Settings_LOCAL Settings_Category = 8
	// Movies and TV category.
	Settings_MOVIES_AND_TV Settings_Category = 9
	// Music and Audio category.
	Settings_MUSIC_AND_AUDIO Settings_Category = 10
	// News category,
	Settings_NEWS Settings_Category = 1
	// Novelty and Humor category.
	Settings_NOVELTY_AND_HUMOR Settings_Category = 11
	// Productivity category.
	Settings_PRODUCTIVITY Settings_Category = 12
	// Shopping category.
	Settings_SHOPPING Settings_Category = 13
	// Social category.
	Settings_SOCIAL Settings_Category = 14
	// Sports category.
	Settings_SPORTS Settings_Category = 15
	// Travel and Transportation category.
	Settings_TRAVEL_AND_TRANSPORTATION Settings_Category = 16
	// Utilities category.
	Settings_UTILITIES Settings_Category = 17
	// Weather category.
	Settings_WEATHER Settings_Category = 18
	// Home Control category.
	Settings_HOME_CONTROL Settings_Category = 19
)

// Enum value maps for Settings_Category.
var (
	Settings_Category_name = map[int32]string{
		0:  "CATEGORY_UNSPECIFIED",
		2:  "BUSINESS_AND_FINANCE",
		3:  "EDUCATION_AND_REFERENCE",
		4:  "FOOD_AND_DRINK",
		5:  "GAMES_AND_TRIVIA",
		6:  "HEALTH_AND_FITNESS",
		20: "KIDS_AND_FAMILY",
		7:  "LIFESTYLE",
		8:  "LOCAL",
		9:  "MOVIES_AND_TV",
		10: "MUSIC_AND_AUDIO",
		1:  "NEWS",
		11: "NOVELTY_AND_HUMOR",
		12: "PRODUCTIVITY",
		13: "SHOPPING",
		14: "SOCIAL",
		15: "SPORTS",
		16: "TRAVEL_AND_TRANSPORTATION",
		17: "UTILITIES",
		18: "WEATHER",
		19: "HOME_CONTROL",
	}
	Settings_Category_value = map[string]int32{
		"CATEGORY_UNSPECIFIED":      0,
		"BUSINESS_AND_FINANCE":      2,
		"EDUCATION_AND_REFERENCE":   3,
		"FOOD_AND_DRINK":            4,
		"GAMES_AND_TRIVIA":          5,
		"HEALTH_AND_FITNESS":        6,
		"KIDS_AND_FAMILY":           20,
		"LIFESTYLE":                 7,
		"LOCAL":                     8,
		"MOVIES_AND_TV":             9,
		"MUSIC_AND_AUDIO":           10,
		"NEWS":                      1,
		"NOVELTY_AND_HUMOR":         11,
		"PRODUCTIVITY":              12,
		"SHOPPING":                  13,
		"SOCIAL":                    14,
		"SPORTS":                    15,
		"TRAVEL_AND_TRANSPORTATION": 16,
		"UTILITIES":                 17,
		"WEATHER":                   18,
		"HOME_CONTROL":              19,
	}
)

func (x Settings_Category) Enum() *Settings_Category {
	p := new(Settings_Category)
	*p = x
	return p
}

func (x Settings_Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Settings_Category) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_settings_proto_enumTypes[0].Descriptor()
}

func (Settings_Category) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_settings_proto_enumTypes[0]
}

func (x Settings_Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Settings_Category.Descriptor instead.
func (Settings_Category) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_settings_proto_rawDescGZIP(), []int{0, 0}
}

// Represents settings of an Actions project that are not locale specific.
type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Actions project id.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Locale which is default for the project. For all files except under
	// `resources/` with no locale in the path, the localized data is attributed
	// to this `default_locale`. For files under `resources/` no locale means that
	// the resource is applicable to all locales.
	DefaultLocale string `protobuf:"bytes,2,opt,name=default_locale,json=defaultLocale,proto3" json:"default_locale,omitempty"`
	// Represents the regions where users can invoke your Actions, which is
	// based on the user's location of presence. Cannot be set if
	// `disabled_regions` is set. If both `enabled_regions` and `disabled_regions`
	// are not specified, users can invoke your Actions in all regions. Each
	// region is represented using the Canonical Name of Adwords geotargets. See
	// https://developers.google.com/adwords/api/docs/appendix/geotargeting
	// Examples include:
	// - "Germany"
	// - "Ghana"
	// - "Greece"
	// - "Grenada"
	// - "United Kingdom"
	// - "United States"
	// - "United States Minor Outlying Islands"
	// - "Uruguay"
	EnabledRegions []string `protobuf:"bytes,3,rep,name=enabled_regions,json=enabledRegions,proto3" json:"enabled_regions,omitempty"`
	// Represents the regions where your Actions are blocked, based on the user's
	// location of presence. Cannot be set if `enabled_regions` is set.
	// Each region is represented using the Canonical Name of Adwords geotargets.
	// See https://developers.google.com/adwords/api/docs/appendix/geotargeting
	// Examples include:
	// - "Germany"
	// - "Ghana"
	// - "Greece"
	// - "Grenada"
	// - "United Kingdom"
	// - "United States"
	// - "United States Minor Outlying Islands"
	// - "Uruguay"
	DisabledRegions []string `protobuf:"bytes,4,rep,name=disabled_regions,json=disabledRegions,proto3" json:"disabled_regions,omitempty"`
	// The category for this Actions project.
	Category Settings_Category `protobuf:"varint,5,opt,name=category,proto3,enum=google.actions.sdk.v2.Settings_Category" json:"category,omitempty"`
	// Whether Actions can use transactions (for example, making
	// reservations, taking orders, etc.). If false, then attempts to use the
	// Transactions APIs fail.
	UsesTransactionsApi bool `protobuf:"varint,6,opt,name=uses_transactions_api,json=usesTransactionsApi,proto3" json:"uses_transactions_api,omitempty"`
	// Whether Actions can perform transactions for digital goods.
	UsesDigitalPurchaseApi bool `protobuf:"varint,7,opt,name=uses_digital_purchase_api,json=usesDigitalPurchaseApi,proto3" json:"uses_digital_purchase_api,omitempty"`
	// Whether Actions use Interactive Canvas.
	UsesInteractiveCanvas bool `protobuf:"varint,8,opt,name=uses_interactive_canvas,json=usesInteractiveCanvas,proto3" json:"uses_interactive_canvas,omitempty"`
	// Whether Actions use the home storage feature.
	UsesHomeStorage bool `protobuf:"varint,17,opt,name=uses_home_storage,json=usesHomeStorage,proto3" json:"uses_home_storage,omitempty"`
	// Whether Actions content is designed for family (DFF).
	DesignedForFamily bool `protobuf:"varint,9,opt,name=designed_for_family,json=designedForFamily,proto3" json:"designed_for_family,omitempty"`
	// Whether Actions contains alcohol or tobacco related content.
	ContainsAlcoholOrTobaccoContent bool `protobuf:"varint,11,opt,name=contains_alcohol_or_tobacco_content,json=containsAlcoholOrTobaccoContent,proto3" json:"contains_alcohol_or_tobacco_content,omitempty"`
	// Whether Actions may leave mic open without an explicit prompt during
	// conversation.
	KeepsMicOpen bool `protobuf:"varint,12,opt,name=keeps_mic_open,json=keepsMicOpen,proto3" json:"keeps_mic_open,omitempty"`
	// The surface requirements that a client surface must support to invoke
	// Actions in this project.
	SurfaceRequirements *SurfaceRequirements `protobuf:"bytes,13,opt,name=surface_requirements,json=surfaceRequirements,proto3" json:"surface_requirements,omitempty"`
	// Free-form testing instructions for Actions reviewer (for example, account
	// linking instructions).
	TestingInstructions string `protobuf:"bytes,14,opt,name=testing_instructions,json=testingInstructions,proto3" json:"testing_instructions,omitempty"`
	// Localized settings for the project's default locale. Every additional
	// locale should have its own settings file in its own directory.
	LocalizedSettings *LocalizedSettings `protobuf:"bytes,15,opt,name=localized_settings,json=localizedSettings,proto3" json:"localized_settings,omitempty"`
	// Allow users to create or link accounts through Google sign-in and/or your
	// own OAuth service.
	AccountLinking *AccountLinking `protobuf:"bytes,16,opt,name=account_linking,json=accountLinking,proto3" json:"account_linking,omitempty"`
	// Android apps selected to acccess Google Play purchases for transactions.
	// This is a selection from the Android apps connected to the actions project
	// to verify brand ownership and enable additional features. See
	// https://developers.google.com/assistant/console/brand-verification for more
	// information.
	SelectedAndroidApps []string `protobuf:"bytes,20,rep,name=selected_android_apps,json=selectedAndroidApps,proto3" json:"selected_android_apps,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_settings_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Settings) GetDefaultLocale() string {
	if x != nil {
		return x.DefaultLocale
	}
	return ""
}

func (x *Settings) GetEnabledRegions() []string {
	if x != nil {
		return x.EnabledRegions
	}
	return nil
}

func (x *Settings) GetDisabledRegions() []string {
	if x != nil {
		return x.DisabledRegions
	}
	return nil
}

func (x *Settings) GetCategory() Settings_Category {
	if x != nil {
		return x.Category
	}
	return Settings_CATEGORY_UNSPECIFIED
}

func (x *Settings) GetUsesTransactionsApi() bool {
	if x != nil {
		return x.UsesTransactionsApi
	}
	return false
}

func (x *Settings) GetUsesDigitalPurchaseApi() bool {
	if x != nil {
		return x.UsesDigitalPurchaseApi
	}
	return false
}

func (x *Settings) GetUsesInteractiveCanvas() bool {
	if x != nil {
		return x.UsesInteractiveCanvas
	}
	return false
}

func (x *Settings) GetUsesHomeStorage() bool {
	if x != nil {
		return x.UsesHomeStorage
	}
	return false
}

func (x *Settings) GetDesignedForFamily() bool {
	if x != nil {
		return x.DesignedForFamily
	}
	return false
}

func (x *Settings) GetContainsAlcoholOrTobaccoContent() bool {
	if x != nil {
		return x.ContainsAlcoholOrTobaccoContent
	}
	return false
}

func (x *Settings) GetKeepsMicOpen() bool {
	if x != nil {
		return x.KeepsMicOpen
	}
	return false
}

func (x *Settings) GetSurfaceRequirements() *SurfaceRequirements {
	if x != nil {
		return x.SurfaceRequirements
	}
	return nil
}

func (x *Settings) GetTestingInstructions() string {
	if x != nil {
		return x.TestingInstructions
	}
	return ""
}

func (x *Settings) GetLocalizedSettings() *LocalizedSettings {
	if x != nil {
		return x.LocalizedSettings
	}
	return nil
}

func (x *Settings) GetAccountLinking() *AccountLinking {
	if x != nil {
		return x.AccountLinking
	}
	return nil
}

func (x *Settings) GetSelectedAndroidApps() []string {
	if x != nil {
		return x.SelectedAndroidApps
	}
	return nil
}

var File_google_actions_sdk_v2_settings_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_settings_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x1a, 0x2b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76,
	0x32, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76,
	0x32, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xed, 0x0a, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x15,
	0x75, 0x73, 0x65, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x75, 0x73, 0x65,
	0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69,
	0x12, 0x39, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x73, 0x65, 0x73, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x41, 0x70, 0x69, 0x12, 0x36, 0x0a, 0x17, 0x75,
	0x73, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x63, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x75, 0x73,
	0x65, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x6e,
	0x76, 0x61, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x68, 0x6f, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x75, 0x73, 0x65, 0x73, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x2e, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x5f,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12,
	0x4c, 0x0a, 0x23, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x5f, 0x61, 0x6c, 0x63, 0x6f,
	0x68, 0x6f, 0x6c, 0x5f, 0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x41, 0x6c, 0x63, 0x6f, 0x68, 0x6f, 0x6c, 0x4f, 0x72, 0x54,
	0x6f, 0x62, 0x61, 0x63, 0x63, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x6b, 0x65, 0x65, 0x70, 0x73, 0x5f, 0x6d, 0x69, 0x63, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6b, 0x65, 0x65, 0x70, 0x73, 0x4d, 0x69, 0x63, 0x4f,
	0x70, 0x65, 0x6e, 0x12, 0x5d, 0x0a, 0x14, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x13, 0x73,
	0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x57, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x11, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4e,
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x32,
	0x0a, 0x15, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x41, 0x70,
	0x70, 0x73, 0x22, 0x9a, 0x03, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x18, 0x0a, 0x14, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x55, 0x53,
	0x49, 0x4e, 0x45, 0x53, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4e, 0x43,
	0x45, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x44, 0x55, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4f, 0x4f, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x44, 0x52, 0x49,
	0x4e, 0x4b, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x41, 0x4d, 0x45, 0x53, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x54, 0x52, 0x49, 0x56, 0x49, 0x41, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x46, 0x49, 0x54, 0x4e, 0x45, 0x53, 0x53,
	0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x49, 0x44, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x46,
	0x41, 0x4d, 0x49, 0x4c, 0x59, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x49, 0x46, 0x45, 0x53,
	0x54, 0x59, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10,
	0x08, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x56, 0x49, 0x45, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x5f,
	0x54, 0x56, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x55, 0x53, 0x49, 0x43, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x57,
	0x53, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x4f, 0x56, 0x45, 0x4c, 0x54, 0x59, 0x5f, 0x41,
	0x4e, 0x44, 0x5f, 0x48, 0x55, 0x4d, 0x4f, 0x52, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x48, 0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f,
	0x43, 0x49, 0x41, 0x4c, 0x10, 0x0e, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x53,
	0x10, 0x0f, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x52, 0x41, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x4e, 0x44,
	0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x10, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x54, 0x49, 0x4c, 0x49, 0x54, 0x49, 0x45, 0x53, 0x10, 0x11,
	0x12, 0x0b, 0x0a, 0x07, 0x57, 0x45, 0x41, 0x54, 0x48, 0x45, 0x52, 0x10, 0x12, 0x12, 0x10, 0x0a,
	0x0c, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x10, 0x13, 0x42,
	0x66, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x42, 0x0d, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x76, 0x32, 0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_settings_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_settings_proto_rawDescData = file_google_actions_sdk_v2_settings_proto_rawDesc
)

func file_google_actions_sdk_v2_settings_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_settings_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_settings_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_settings_proto_rawDescData
}

var file_google_actions_sdk_v2_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_actions_sdk_v2_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_settings_proto_goTypes = []interface{}{
	(Settings_Category)(0),      // 0: google.actions.sdk.v2.Settings.Category
	(*Settings)(nil),            // 1: google.actions.sdk.v2.Settings
	(*SurfaceRequirements)(nil), // 2: google.actions.sdk.v2.SurfaceRequirements
	(*LocalizedSettings)(nil),   // 3: google.actions.sdk.v2.LocalizedSettings
	(*AccountLinking)(nil),      // 4: google.actions.sdk.v2.AccountLinking
}
var file_google_actions_sdk_v2_settings_proto_depIdxs = []int32{
	0, // 0: google.actions.sdk.v2.Settings.category:type_name -> google.actions.sdk.v2.Settings.Category
	2, // 1: google.actions.sdk.v2.Settings.surface_requirements:type_name -> google.actions.sdk.v2.SurfaceRequirements
	3, // 2: google.actions.sdk.v2.Settings.localized_settings:type_name -> google.actions.sdk.v2.LocalizedSettings
	4, // 3: google.actions.sdk.v2.Settings.account_linking:type_name -> google.actions.sdk.v2.AccountLinking
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_settings_proto_init() }
func file_google_actions_sdk_v2_settings_proto_init() {
	if File_google_actions_sdk_v2_settings_proto != nil {
		return
	}
	file_google_actions_sdk_v2_account_linking_proto_init()
	file_google_actions_sdk_v2_localized_settings_proto_init()
	file_google_actions_sdk_v2_surface_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_settings_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_settings_proto_depIdxs,
		EnumInfos:         file_google_actions_sdk_v2_settings_proto_enumTypes,
		MessageInfos:      file_google_actions_sdk_v2_settings_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_settings_proto = out.File
	file_google_actions_sdk_v2_settings_proto_rawDesc = nil
	file_google_actions_sdk_v2_settings_proto_goTypes = nil
	file_google_actions_sdk_v2_settings_proto_depIdxs = nil
}
