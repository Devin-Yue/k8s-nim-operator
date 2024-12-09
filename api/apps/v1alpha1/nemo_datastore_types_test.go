/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"
	"testing"

	rendertypes "github.com/NVIDIA/k8s-nim-operator/internal/render/types"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNemoDatastore_GetPVCName(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	type args struct {
		pvc PersistentVolumeClaim
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "pvc name is empty",
			fields: fields{
				ObjectMeta: metav1.ObjectMeta{Name: "nemo-datastore"},
			},
			args: args{
				pvc: PersistentVolumeClaim{},
			},
			want: "nemo-datastore-pvc",
		},
		{
			name: "pvc name is not empty",
			fields: fields{
				ObjectMeta: metav1.ObjectMeta{Name: "nemo-datastore"},
			},
			args: args{
				pvc: PersistentVolumeClaim{Name: "custom-pvc-name"},
			},
			want: "custom-pvc-name",
		},
		{
			name: "pvc name is special characters",
			fields: fields{
				ObjectMeta: metav1.ObjectMeta{Name: "nemo-datastore"},
			},
			args: args{
				pvc: PersistentVolumeClaim{Name: "custom-!@#$%^&*()_+"},
			},
			want: "custom-!@#$%^&*()_+",
		},
		{
			name: "pvc name is space",
			fields: fields{
				ObjectMeta: metav1.ObjectMeta{Name: "nemo-datastore"},
			},
			args: args{
				pvc: PersistentVolumeClaim{Name: " "},
			},
			want: " ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetPVCName(tt.args.pvc); got != tt.want {
				t.Errorf("NemoDatastore.GetPVCName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetStandardSelectorLabels(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetStandardSelectorLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetStandardSelectorLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetStandardLabels(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetStandardLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetStandardLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetStandardEnv(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []corev1.EnvVar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetStandardEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetStandardEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetEnvFrom(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []corev1.EnvFromSource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetEnvFrom(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetEnvFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetInitContainers(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []corev1.Container
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetInitContainers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetInitContainers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetStandardAnnotations(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetStandardAnnotations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetStandardAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetNemoDatastoreAnnotations(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetNemoDatastoreAnnotations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetNemoDatastoreAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceLabels(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetServiceLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetSelectorLabels(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetSelectorLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetSelectorLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetNodeSelector(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetNodeSelector(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetNodeSelector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetTolerations(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []corev1.Toleration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetTolerations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetTolerations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetPodAffinity(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *corev1.PodAffinity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetPodAffinity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetPodAffinity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetContainerName(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetContainerName(); got != tt.want {
				t.Errorf("NemoDatastore.GetContainerName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetCommand(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetCommand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetArgs(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetEnv(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []corev1.EnvVar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetImage(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetImage(); got != tt.want {
				t.Errorf("NemoDatastore.GetImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetImagePullSecrets(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetImagePullSecrets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetImagePullSecrets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetImagePullPolicy(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetImagePullPolicy(); got != tt.want {
				t.Errorf("NemoDatastore.GetImagePullPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetResources(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *corev1.ResourceRequirements
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetResources(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetResources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetLivenessProbe(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *corev1.Probe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetLivenessProbe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetLivenessProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetDefaultLivenessProbe(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *corev1.Probe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetDefaultLivenessProbe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetDefaultLivenessProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetReadinessProbe(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *corev1.Probe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetReadinessProbe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetReadinessProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetDefaultReadinessProbe(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *corev1.Probe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetDefaultReadinessProbe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetDefaultReadinessProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetStartupProbe(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *corev1.Probe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetStartupProbe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetStartupProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetDefaultStartupProbe(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *corev1.Probe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetDefaultStartupProbe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetDefaultStartupProbe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetVolumes(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []corev1.Volume
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetVolumes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetVolumes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetVolumeMounts(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   []corev1.VolumeMount
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetVolumeMounts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetVolumeMounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceAccountName(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceAccountName(); got != tt.want {
				t.Errorf("NemoDatastore.GetServiceAccountName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetRuntimeClass(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetRuntimeClass(); got != tt.want {
				t.Errorf("NemoDatastore.GetRuntimeClass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetHPA(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   HorizontalPodAutoscalerSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetHPA(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetHPA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceMonitor(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   ServiceMonitor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceMonitor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetServiceMonitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetReplicas(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetReplicas(); got != tt.want {
				t.Errorf("NemoDatastore.GetReplicas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetDeploymentKind(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetDeploymentKind(); got != tt.want {
				t.Errorf("NemoDatastore.GetDeploymentKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_IsAutoScalingEnabled(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.IsAutoScalingEnabled(); got != tt.want {
				t.Errorf("NemoDatastore.IsAutoScalingEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_IsIngressEnabled(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.IsIngressEnabled(); got != tt.want {
				t.Errorf("NemoDatastore.IsIngressEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetIngressSpec(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   networkingv1.IngressSpec
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetIngressSpec(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetIngressSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_IsServiceMonitorEnabled(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.IsServiceMonitorEnabled(); got != tt.want {
				t.Errorf("NemoDatastore.IsServiceMonitorEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServicePort(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServicePort(); got != tt.want {
				t.Errorf("NemoDatastore.GetServicePort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceType(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceType(); got != tt.want {
				t.Errorf("NemoDatastore.GetServiceType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetUserID(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetUserID(); got != tt.want {
				t.Errorf("NemoDatastore.GetUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetGroupID(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetGroupID(); got != tt.want {
				t.Errorf("NemoDatastore.GetGroupID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceAccountParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.ServiceAccountParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceAccountParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetServiceAccountParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetDeploymentParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.DeploymentParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetDeploymentParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetDeploymentParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetStatefulSetParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.StatefulSetParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetStatefulSetParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetStatefulSetParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.ServiceParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetServiceParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetIngressParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.IngressParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetIngressParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetIngressParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetRoleParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.RoleParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetRoleParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetRoleParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetRoleBindingParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.RoleBindingParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetRoleBindingParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetRoleBindingParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetHPAParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.HPAParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetHPAParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetHPAParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetSCCParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.SCCParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetSCCParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetSCCParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceMonitorParams(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   *rendertypes.ServiceMonitorParams
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceMonitorParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetServiceMonitorParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetIngressAnnotations(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetIngressAnnotations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetIngressAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceAnnotations(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceAnnotations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetServiceAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetHPAAnnotations(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetHPAAnnotations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetHPAAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNemoDatastore_GetServiceMonitorAnnotations(t *testing.T) {
	type fields struct {
		TypeMeta   metav1.TypeMeta
		ObjectMeta metav1.ObjectMeta
		Spec       NemoDatastoreSpec
		Status     NemoDatastoreStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NemoDatastore{
				TypeMeta:   tt.fields.TypeMeta,
				ObjectMeta: tt.fields.ObjectMeta,
				Spec:       tt.fields.Spec,
				Status:     tt.fields.Status,
			}
			if got := n.GetServiceMonitorAnnotations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NemoDatastore.GetServiceMonitorAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}
