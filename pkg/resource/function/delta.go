// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package function

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customPreCompare(delta, a, b)

	if len(a.ko.Spec.Architectures) != len(b.ko.Spec.Architectures) {
		delta.Add("Spec.Architectures", a.ko.Spec.Architectures, b.ko.Spec.Architectures)
	} else if len(a.ko.Spec.Architectures) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.Architectures, b.ko.Spec.Architectures) {
			delta.Add("Spec.Architectures", a.ko.Spec.Architectures, b.ko.Spec.Architectures)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CodeSigningConfigARN, b.ko.Spec.CodeSigningConfigARN) {
		delta.Add("Spec.CodeSigningConfigARN", a.ko.Spec.CodeSigningConfigARN, b.ko.Spec.CodeSigningConfigARN)
	} else if a.ko.Spec.CodeSigningConfigARN != nil && b.ko.Spec.CodeSigningConfigARN != nil {
		if *a.ko.Spec.CodeSigningConfigARN != *b.ko.Spec.CodeSigningConfigARN {
			delta.Add("Spec.CodeSigningConfigARN", a.ko.Spec.CodeSigningConfigARN, b.ko.Spec.CodeSigningConfigARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DeadLetterConfig, b.ko.Spec.DeadLetterConfig) {
		delta.Add("Spec.DeadLetterConfig", a.ko.Spec.DeadLetterConfig, b.ko.Spec.DeadLetterConfig)
	} else if a.ko.Spec.DeadLetterConfig != nil && b.ko.Spec.DeadLetterConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.DeadLetterConfig.TargetARN, b.ko.Spec.DeadLetterConfig.TargetARN) {
			delta.Add("Spec.DeadLetterConfig.TargetARN", a.ko.Spec.DeadLetterConfig.TargetARN, b.ko.Spec.DeadLetterConfig.TargetARN)
		} else if a.ko.Spec.DeadLetterConfig.TargetARN != nil && b.ko.Spec.DeadLetterConfig.TargetARN != nil {
			if *a.ko.Spec.DeadLetterConfig.TargetARN != *b.ko.Spec.DeadLetterConfig.TargetARN {
				delta.Add("Spec.DeadLetterConfig.TargetARN", a.ko.Spec.DeadLetterConfig.TargetARN, b.ko.Spec.DeadLetterConfig.TargetARN)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Environment, b.ko.Spec.Environment) {
		delta.Add("Spec.Environment", a.ko.Spec.Environment, b.ko.Spec.Environment)
	} else if a.ko.Spec.Environment != nil && b.ko.Spec.Environment != nil {
		if len(a.ko.Spec.Environment.Variables) != len(b.ko.Spec.Environment.Variables) {
			delta.Add("Spec.Environment.Variables", a.ko.Spec.Environment.Variables, b.ko.Spec.Environment.Variables)
		} else if len(a.ko.Spec.Environment.Variables) > 0 {
			if !ackcompare.MapStringStringPEqual(a.ko.Spec.Environment.Variables, b.ko.Spec.Environment.Variables) {
				delta.Add("Spec.Environment.Variables", a.ko.Spec.Environment.Variables, b.ko.Spec.Environment.Variables)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EphemeralStorage, b.ko.Spec.EphemeralStorage) {
		delta.Add("Spec.EphemeralStorage", a.ko.Spec.EphemeralStorage, b.ko.Spec.EphemeralStorage)
	} else if a.ko.Spec.EphemeralStorage != nil && b.ko.Spec.EphemeralStorage != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.EphemeralStorage.Size, b.ko.Spec.EphemeralStorage.Size) {
			delta.Add("Spec.EphemeralStorage.Size", a.ko.Spec.EphemeralStorage.Size, b.ko.Spec.EphemeralStorage.Size)
		} else if a.ko.Spec.EphemeralStorage.Size != nil && b.ko.Spec.EphemeralStorage.Size != nil {
			if *a.ko.Spec.EphemeralStorage.Size != *b.ko.Spec.EphemeralStorage.Size {
				delta.Add("Spec.EphemeralStorage.Size", a.ko.Spec.EphemeralStorage.Size, b.ko.Spec.EphemeralStorage.Size)
			}
		}
	}
	if len(a.ko.Spec.FileSystemConfigs) != len(b.ko.Spec.FileSystemConfigs) {
		delta.Add("Spec.FileSystemConfigs", a.ko.Spec.FileSystemConfigs, b.ko.Spec.FileSystemConfigs)
	} else if len(a.ko.Spec.FileSystemConfigs) > 0 {
		if !reflect.DeepEqual(a.ko.Spec.FileSystemConfigs, b.ko.Spec.FileSystemConfigs) {
			delta.Add("Spec.FileSystemConfigs", a.ko.Spec.FileSystemConfigs, b.ko.Spec.FileSystemConfigs)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig, b.ko.Spec.FunctionEventInvokeConfig) {
		delta.Add("Spec.FunctionEventInvokeConfig", a.ko.Spec.FunctionEventInvokeConfig, b.ko.Spec.FunctionEventInvokeConfig)
	} else if a.ko.Spec.FunctionEventInvokeConfig != nil && b.ko.Spec.FunctionEventInvokeConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig) {
			delta.Add("Spec.FunctionEventInvokeConfig.DestinationConfig", a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig)
		} else if a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig != nil && b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure) {
				delta.Add("Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure", a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure)
			} else if a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure != nil && b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination) {
					delta.Add("Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination", a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination)
				} else if a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination != nil && b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination != nil {
					if *a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination != *b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination {
						delta.Add("Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination", a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnFailure.Destination)
					}
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess) {
				delta.Add("Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess", a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess)
			} else if a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess != nil && b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination) {
					delta.Add("Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination", a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination)
				} else if a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination != nil && b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination != nil {
					if *a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination != *b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination {
						delta.Add("Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination", a.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination, b.ko.Spec.FunctionEventInvokeConfig.DestinationConfig.OnSuccess.Destination)
					}
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.FunctionName, b.ko.Spec.FunctionEventInvokeConfig.FunctionName) {
			delta.Add("Spec.FunctionEventInvokeConfig.FunctionName", a.ko.Spec.FunctionEventInvokeConfig.FunctionName, b.ko.Spec.FunctionEventInvokeConfig.FunctionName)
		} else if a.ko.Spec.FunctionEventInvokeConfig.FunctionName != nil && b.ko.Spec.FunctionEventInvokeConfig.FunctionName != nil {
			if *a.ko.Spec.FunctionEventInvokeConfig.FunctionName != *b.ko.Spec.FunctionEventInvokeConfig.FunctionName {
				delta.Add("Spec.FunctionEventInvokeConfig.FunctionName", a.ko.Spec.FunctionEventInvokeConfig.FunctionName, b.ko.Spec.FunctionEventInvokeConfig.FunctionName)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds, b.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds) {
			delta.Add("Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds", a.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds, b.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds)
		} else if a.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds != nil && b.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds != nil {
			if *a.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds != *b.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds {
				delta.Add("Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds", a.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds, b.ko.Spec.FunctionEventInvokeConfig.MaximumEventAgeInSeconds)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts, b.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts) {
			delta.Add("Spec.FunctionEventInvokeConfig.MaximumRetryAttempts", a.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts, b.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts)
		} else if a.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts != nil && b.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts != nil {
			if *a.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts != *b.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts {
				delta.Add("Spec.FunctionEventInvokeConfig.MaximumRetryAttempts", a.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts, b.ko.Spec.FunctionEventInvokeConfig.MaximumRetryAttempts)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.FunctionEventInvokeConfig.Qualifier, b.ko.Spec.FunctionEventInvokeConfig.Qualifier) {
			delta.Add("Spec.FunctionEventInvokeConfig.Qualifier", a.ko.Spec.FunctionEventInvokeConfig.Qualifier, b.ko.Spec.FunctionEventInvokeConfig.Qualifier)
		} else if a.ko.Spec.FunctionEventInvokeConfig.Qualifier != nil && b.ko.Spec.FunctionEventInvokeConfig.Qualifier != nil {
			if *a.ko.Spec.FunctionEventInvokeConfig.Qualifier != *b.ko.Spec.FunctionEventInvokeConfig.Qualifier {
				delta.Add("Spec.FunctionEventInvokeConfig.Qualifier", a.ko.Spec.FunctionEventInvokeConfig.Qualifier, b.ko.Spec.FunctionEventInvokeConfig.Qualifier)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Handler, b.ko.Spec.Handler) {
		delta.Add("Spec.Handler", a.ko.Spec.Handler, b.ko.Spec.Handler)
	} else if a.ko.Spec.Handler != nil && b.ko.Spec.Handler != nil {
		if *a.ko.Spec.Handler != *b.ko.Spec.Handler {
			delta.Add("Spec.Handler", a.ko.Spec.Handler, b.ko.Spec.Handler)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ImageConfig, b.ko.Spec.ImageConfig) {
		delta.Add("Spec.ImageConfig", a.ko.Spec.ImageConfig, b.ko.Spec.ImageConfig)
	} else if a.ko.Spec.ImageConfig != nil && b.ko.Spec.ImageConfig != nil {
		if len(a.ko.Spec.ImageConfig.Command) != len(b.ko.Spec.ImageConfig.Command) {
			delta.Add("Spec.ImageConfig.Command", a.ko.Spec.ImageConfig.Command, b.ko.Spec.ImageConfig.Command)
		} else if len(a.ko.Spec.ImageConfig.Command) > 0 {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.ImageConfig.Command, b.ko.Spec.ImageConfig.Command) {
				delta.Add("Spec.ImageConfig.Command", a.ko.Spec.ImageConfig.Command, b.ko.Spec.ImageConfig.Command)
			}
		}
		if len(a.ko.Spec.ImageConfig.EntryPoint) != len(b.ko.Spec.ImageConfig.EntryPoint) {
			delta.Add("Spec.ImageConfig.EntryPoint", a.ko.Spec.ImageConfig.EntryPoint, b.ko.Spec.ImageConfig.EntryPoint)
		} else if len(a.ko.Spec.ImageConfig.EntryPoint) > 0 {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.ImageConfig.EntryPoint, b.ko.Spec.ImageConfig.EntryPoint) {
				delta.Add("Spec.ImageConfig.EntryPoint", a.ko.Spec.ImageConfig.EntryPoint, b.ko.Spec.ImageConfig.EntryPoint)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ImageConfig.WorkingDirectory, b.ko.Spec.ImageConfig.WorkingDirectory) {
			delta.Add("Spec.ImageConfig.WorkingDirectory", a.ko.Spec.ImageConfig.WorkingDirectory, b.ko.Spec.ImageConfig.WorkingDirectory)
		} else if a.ko.Spec.ImageConfig.WorkingDirectory != nil && b.ko.Spec.ImageConfig.WorkingDirectory != nil {
			if *a.ko.Spec.ImageConfig.WorkingDirectory != *b.ko.Spec.ImageConfig.WorkingDirectory {
				delta.Add("Spec.ImageConfig.WorkingDirectory", a.ko.Spec.ImageConfig.WorkingDirectory, b.ko.Spec.ImageConfig.WorkingDirectory)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSKeyARN, b.ko.Spec.KMSKeyARN) {
		delta.Add("Spec.KMSKeyARN", a.ko.Spec.KMSKeyARN, b.ko.Spec.KMSKeyARN)
	} else if a.ko.Spec.KMSKeyARN != nil && b.ko.Spec.KMSKeyARN != nil {
		if *a.ko.Spec.KMSKeyARN != *b.ko.Spec.KMSKeyARN {
			delta.Add("Spec.KMSKeyARN", a.ko.Spec.KMSKeyARN, b.ko.Spec.KMSKeyARN)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.KMSKeyRef, b.ko.Spec.KMSKeyRef) {
		delta.Add("Spec.KMSKeyRef", a.ko.Spec.KMSKeyRef, b.ko.Spec.KMSKeyRef)
	}
	if len(a.ko.Spec.Layers) != len(b.ko.Spec.Layers) {
		delta.Add("Spec.Layers", a.ko.Spec.Layers, b.ko.Spec.Layers)
	} else if len(a.ko.Spec.Layers) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.Layers, b.ko.Spec.Layers) {
			delta.Add("Spec.Layers", a.ko.Spec.Layers, b.ko.Spec.Layers)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.MemorySize, b.ko.Spec.MemorySize) {
		delta.Add("Spec.MemorySize", a.ko.Spec.MemorySize, b.ko.Spec.MemorySize)
	} else if a.ko.Spec.MemorySize != nil && b.ko.Spec.MemorySize != nil {
		if *a.ko.Spec.MemorySize != *b.ko.Spec.MemorySize {
			delta.Add("Spec.MemorySize", a.ko.Spec.MemorySize, b.ko.Spec.MemorySize)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PackageType, b.ko.Spec.PackageType) {
		delta.Add("Spec.PackageType", a.ko.Spec.PackageType, b.ko.Spec.PackageType)
	} else if a.ko.Spec.PackageType != nil && b.ko.Spec.PackageType != nil {
		if *a.ko.Spec.PackageType != *b.ko.Spec.PackageType {
			delta.Add("Spec.PackageType", a.ko.Spec.PackageType, b.ko.Spec.PackageType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Publish, b.ko.Spec.Publish) {
		delta.Add("Spec.Publish", a.ko.Spec.Publish, b.ko.Spec.Publish)
	} else if a.ko.Spec.Publish != nil && b.ko.Spec.Publish != nil {
		if *a.ko.Spec.Publish != *b.ko.Spec.Publish {
			delta.Add("Spec.Publish", a.ko.Spec.Publish, b.ko.Spec.Publish)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ReservedConcurrentExecutions, b.ko.Spec.ReservedConcurrentExecutions) {
		delta.Add("Spec.ReservedConcurrentExecutions", a.ko.Spec.ReservedConcurrentExecutions, b.ko.Spec.ReservedConcurrentExecutions)
	} else if a.ko.Spec.ReservedConcurrentExecutions != nil && b.ko.Spec.ReservedConcurrentExecutions != nil {
		if *a.ko.Spec.ReservedConcurrentExecutions != *b.ko.Spec.ReservedConcurrentExecutions {
			delta.Add("Spec.ReservedConcurrentExecutions", a.ko.Spec.ReservedConcurrentExecutions, b.ko.Spec.ReservedConcurrentExecutions)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Role, b.ko.Spec.Role) {
		delta.Add("Spec.Role", a.ko.Spec.Role, b.ko.Spec.Role)
	} else if a.ko.Spec.Role != nil && b.ko.Spec.Role != nil {
		if *a.ko.Spec.Role != *b.ko.Spec.Role {
			delta.Add("Spec.Role", a.ko.Spec.Role, b.ko.Spec.Role)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.RoleRef, b.ko.Spec.RoleRef) {
		delta.Add("Spec.RoleRef", a.ko.Spec.RoleRef, b.ko.Spec.RoleRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Runtime, b.ko.Spec.Runtime) {
		delta.Add("Spec.Runtime", a.ko.Spec.Runtime, b.ko.Spec.Runtime)
	} else if a.ko.Spec.Runtime != nil && b.ko.Spec.Runtime != nil {
		if *a.ko.Spec.Runtime != *b.ko.Spec.Runtime {
			delta.Add("Spec.Runtime", a.ko.Spec.Runtime, b.ko.Spec.Runtime)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SnapStart, b.ko.Spec.SnapStart) {
		delta.Add("Spec.SnapStart", a.ko.Spec.SnapStart, b.ko.Spec.SnapStart)
	} else if a.ko.Spec.SnapStart != nil && b.ko.Spec.SnapStart != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.SnapStart.ApplyOn, b.ko.Spec.SnapStart.ApplyOn) {
			delta.Add("Spec.SnapStart.ApplyOn", a.ko.Spec.SnapStart.ApplyOn, b.ko.Spec.SnapStart.ApplyOn)
		} else if a.ko.Spec.SnapStart.ApplyOn != nil && b.ko.Spec.SnapStart.ApplyOn != nil {
			if *a.ko.Spec.SnapStart.ApplyOn != *b.ko.Spec.SnapStart.ApplyOn {
				delta.Add("Spec.SnapStart.ApplyOn", a.ko.Spec.SnapStart.ApplyOn, b.ko.Spec.SnapStart.ApplyOn)
			}
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Timeout, b.ko.Spec.Timeout) {
		delta.Add("Spec.Timeout", a.ko.Spec.Timeout, b.ko.Spec.Timeout)
	} else if a.ko.Spec.Timeout != nil && b.ko.Spec.Timeout != nil {
		if *a.ko.Spec.Timeout != *b.ko.Spec.Timeout {
			delta.Add("Spec.Timeout", a.ko.Spec.Timeout, b.ko.Spec.Timeout)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TracingConfig, b.ko.Spec.TracingConfig) {
		delta.Add("Spec.TracingConfig", a.ko.Spec.TracingConfig, b.ko.Spec.TracingConfig)
	} else if a.ko.Spec.TracingConfig != nil && b.ko.Spec.TracingConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.TracingConfig.Mode, b.ko.Spec.TracingConfig.Mode) {
			delta.Add("Spec.TracingConfig.Mode", a.ko.Spec.TracingConfig.Mode, b.ko.Spec.TracingConfig.Mode)
		} else if a.ko.Spec.TracingConfig.Mode != nil && b.ko.Spec.TracingConfig.Mode != nil {
			if *a.ko.Spec.TracingConfig.Mode != *b.ko.Spec.TracingConfig.Mode {
				delta.Add("Spec.TracingConfig.Mode", a.ko.Spec.TracingConfig.Mode, b.ko.Spec.TracingConfig.Mode)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.VPCConfig, b.ko.Spec.VPCConfig) {
		delta.Add("Spec.VPCConfig", a.ko.Spec.VPCConfig, b.ko.Spec.VPCConfig)
	} else if a.ko.Spec.VPCConfig != nil && b.ko.Spec.VPCConfig != nil {
		if len(a.ko.Spec.VPCConfig.SecurityGroupIDs) != len(b.ko.Spec.VPCConfig.SecurityGroupIDs) {
			delta.Add("Spec.VPCConfig.SecurityGroupIDs", a.ko.Spec.VPCConfig.SecurityGroupIDs, b.ko.Spec.VPCConfig.SecurityGroupIDs)
		} else if len(a.ko.Spec.VPCConfig.SecurityGroupIDs) > 0 {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCConfig.SecurityGroupIDs, b.ko.Spec.VPCConfig.SecurityGroupIDs) {
				delta.Add("Spec.VPCConfig.SecurityGroupIDs", a.ko.Spec.VPCConfig.SecurityGroupIDs, b.ko.Spec.VPCConfig.SecurityGroupIDs)
			}
		}
		if len(a.ko.Spec.VPCConfig.SubnetIDs) != len(b.ko.Spec.VPCConfig.SubnetIDs) {
			delta.Add("Spec.VPCConfig.SubnetIDs", a.ko.Spec.VPCConfig.SubnetIDs, b.ko.Spec.VPCConfig.SubnetIDs)
		} else if len(a.ko.Spec.VPCConfig.SubnetIDs) > 0 {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCConfig.SubnetIDs, b.ko.Spec.VPCConfig.SubnetIDs) {
				delta.Add("Spec.VPCConfig.SubnetIDs", a.ko.Spec.VPCConfig.SubnetIDs, b.ko.Spec.VPCConfig.SubnetIDs)
			}
		}
	}

	return delta
}
