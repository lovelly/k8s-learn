// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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
// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	learn "k8s-learn/apiserver/pkg/apis/learn"
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Nginx)(nil), (*learn.Nginx)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Nginx_To_learn_Nginx(a.(*Nginx), b.(*learn.Nginx), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*learn.Nginx)(nil), (*Nginx)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_learn_Nginx_To_v1_Nginx(a.(*learn.Nginx), b.(*Nginx), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NginxList)(nil), (*learn.NginxList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NginxList_To_learn_NginxList(a.(*NginxList), b.(*learn.NginxList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*learn.NginxList)(nil), (*NginxList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_learn_NginxList_To_v1_NginxList(a.(*learn.NginxList), b.(*NginxList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NginxSpec)(nil), (*learn.NginxSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NginxSpec_To_learn_NginxSpec(a.(*NginxSpec), b.(*learn.NginxSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*learn.NginxSpec)(nil), (*NginxSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_learn_NginxSpec_To_v1_NginxSpec(a.(*learn.NginxSpec), b.(*NginxSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NginxStatus)(nil), (*learn.NginxStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NginxStatus_To_learn_NginxStatus(a.(*NginxStatus), b.(*learn.NginxStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*learn.NginxStatus)(nil), (*NginxStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_learn_NginxStatus_To_v1_NginxStatus(a.(*learn.NginxStatus), b.(*NginxStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Nginx_To_learn_Nginx(in *Nginx, out *learn.Nginx, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_NginxSpec_To_learn_NginxSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_NginxStatus_To_learn_NginxStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Nginx_To_learn_Nginx is an autogenerated conversion function.
func Convert_v1_Nginx_To_learn_Nginx(in *Nginx, out *learn.Nginx, s conversion.Scope) error {
	return autoConvert_v1_Nginx_To_learn_Nginx(in, out, s)
}

func autoConvert_learn_Nginx_To_v1_Nginx(in *learn.Nginx, out *Nginx, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_learn_NginxSpec_To_v1_NginxSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_learn_NginxStatus_To_v1_NginxStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_learn_Nginx_To_v1_Nginx is an autogenerated conversion function.
func Convert_learn_Nginx_To_v1_Nginx(in *learn.Nginx, out *Nginx, s conversion.Scope) error {
	return autoConvert_learn_Nginx_To_v1_Nginx(in, out, s)
}

func autoConvert_v1_NginxList_To_learn_NginxList(in *NginxList, out *learn.NginxList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]learn.Nginx)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_NginxList_To_learn_NginxList is an autogenerated conversion function.
func Convert_v1_NginxList_To_learn_NginxList(in *NginxList, out *learn.NginxList, s conversion.Scope) error {
	return autoConvert_v1_NginxList_To_learn_NginxList(in, out, s)
}

func autoConvert_learn_NginxList_To_v1_NginxList(in *learn.NginxList, out *NginxList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Nginx)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_learn_NginxList_To_v1_NginxList is an autogenerated conversion function.
func Convert_learn_NginxList_To_v1_NginxList(in *learn.NginxList, out *NginxList, s conversion.Scope) error {
	return autoConvert_learn_NginxList_To_v1_NginxList(in, out, s)
}

func autoConvert_v1_NginxSpec_To_learn_NginxSpec(in *NginxSpec, out *learn.NginxSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1_NginxSpec_To_learn_NginxSpec is an autogenerated conversion function.
func Convert_v1_NginxSpec_To_learn_NginxSpec(in *NginxSpec, out *learn.NginxSpec, s conversion.Scope) error {
	return autoConvert_v1_NginxSpec_To_learn_NginxSpec(in, out, s)
}

func autoConvert_learn_NginxSpec_To_v1_NginxSpec(in *learn.NginxSpec, out *NginxSpec, s conversion.Scope) error {
	return nil
}

// Convert_learn_NginxSpec_To_v1_NginxSpec is an autogenerated conversion function.
func Convert_learn_NginxSpec_To_v1_NginxSpec(in *learn.NginxSpec, out *NginxSpec, s conversion.Scope) error {
	return autoConvert_learn_NginxSpec_To_v1_NginxSpec(in, out, s)
}

func autoConvert_v1_NginxStatus_To_learn_NginxStatus(in *NginxStatus, out *learn.NginxStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1_NginxStatus_To_learn_NginxStatus is an autogenerated conversion function.
func Convert_v1_NginxStatus_To_learn_NginxStatus(in *NginxStatus, out *learn.NginxStatus, s conversion.Scope) error {
	return autoConvert_v1_NginxStatus_To_learn_NginxStatus(in, out, s)
}

func autoConvert_learn_NginxStatus_To_v1_NginxStatus(in *learn.NginxStatus, out *NginxStatus, s conversion.Scope) error {
	return nil
}

// Convert_learn_NginxStatus_To_v1_NginxStatus is an autogenerated conversion function.
func Convert_learn_NginxStatus_To_v1_NginxStatus(in *learn.NginxStatus, out *NginxStatus, s conversion.Scope) error {
	return autoConvert_learn_NginxStatus_To_v1_NginxStatus(in, out, s)
}