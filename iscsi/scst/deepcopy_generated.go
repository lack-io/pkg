// +build !ignore_autogenerated

// Copyright 2021 lack
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
// Code generated by deepcopy-gen. Do NOT EDIT.

package scst

// DeepCopyInto is an auto-generated deepcopy function, coping the receiver, writing into out. in must be no-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	return
}

// DeepCopy is an auto-generated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an auto-generated deepcopy function, coping the receiver, writing into out. in must be no-nil.
func (in *Driver) DeepCopyInto(out *Driver) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make(map[string]*Target, len(*in))
		for key, val := range *in {
			var outVal *Target
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Target)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an auto-generated deepcopy function, copying the receiver, creating a new Driver.
func (in *Driver) DeepCopy() *Driver {
	if in == nil {
		return nil
	}
	out := new(Driver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an auto-generated deepcopy function, coping the receiver, writing into out. in must be no-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	if in.Luns != nil {
		in, out := &in.Luns, &out.Luns
		*out = make([]*Lun, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*out)[i], &(*out)[i]
				*out = new(Lun)
				**out = **in
			}
		}
	}
	if in.Initiators != nil {
		in, out := &in.Initiators, &out.Initiators
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an auto-generated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an auto-generated deepcopy function, coping the receiver, writing into out. in must be no-nil.
func (in *Handler) DeepCopyInto(out *Handler) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make(map[string]*Device, len(*in))
		for key, val := range *in {
			var outVal *Device
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Device)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an auto-generated deepcopy function, copying the receiver, creating a new Handler.
func (in *Handler) DeepCopy() *Handler {
	if in == nil {
		return nil
	}
	out := new(Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an auto-generated deepcopy function, coping the receiver, writing into out. in must be no-nil.
func (in *Lun) DeepCopyInto(out *Lun) {
	*out = *in
	return
}

// DeepCopy is an auto-generated deepcopy function, copying the receiver, creating a new Lun.
func (in *Lun) DeepCopy() *Lun {
	if in == nil {
		return nil
	}
	out := new(Lun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an auto-generated deepcopy function, coping the receiver, writing into out. in must be no-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make(map[string]*Group, len(*in))
		for key, val := range *in {
			var outVal *Group
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Group)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Luns != nil {
		in, out := &in.Luns, &out.Luns
		*out = make([]*Lun, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*out)[i], &(*out)[i]
				*out = new(Lun)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an auto-generated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}