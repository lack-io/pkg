// Copyright 2021 lack
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

package scst

var ScstTmpl = `# Automatically generated by SCST Configurator {{.Version}}.

{{ range $handler := .Handlers }}
HANDLER {{ $handler.Name }} {
  {{- range $device := $handler.Devices }}
  DEVICE {{ $device.Name }} {
    filename {{ $device.Filename }}
    size {{ $device.Size }}
  } 
  {{ end }}
}
{{ end }} 
{{ range $driver := .Drivers }}
TARGET_DRIVER {{ $driver.Name }} { {{- $name := "iscsi" -}} {{ if eq $driver.Name $name }}
  enable {{ $driver.Enabled }} {{ end }} 
{{- range $target := $driver.Targets }}
  TARGET {{ $target.Name }} { {{ if eq $driver.Name $name }}
    enabled {{ $target.Enabled }}
    rel_tgt_id {{ $target.Id }} {{ end }}
    {{ range $group := $target.Groups }}
    GROUP {{ $group.Name }} { {{ range $lun := $group.Luns }}
      LUN {{ $lun.Id }} {{ $lun.Device }}{{ end }}
      {{- range $initiator := $group.Initiators }} 
      INITIATOR {{ $initiator }} {{ end }}
    }{{ end }} {{- range $lun := $target.Luns }}
    LUN {{ $lun.Id }} {{ $lun.Device }} {{ end }}
  }
  {{ end }}
}
{{ end }}
`