package main

import (
	"bytes"
	"text/template"
)

type DRAC struct {
	Host     string
	Username string
	Password string
}

func (d *DRAC) Viewer() (string, error) {

	// Generate a JNLP viewer from the template
	// Injecting the host/user/pass information
	buff := bytes.NewBufferString("")
	err := template.Must(template.New("viewer").Parse(viewer)).Execute(buff, d)
	return buff.String(), err

}

const viewer string = `
<?xml version="1.0" encoding="UTF-8"?>
<jnlp codebase="https://{{ .Host }}:443" spec="1.0+">
<information>
  <title>iDRAC7 Virtual Console Client</title>
  <vendor>Dell Inc.</vendor>
   <icon href="https://{{ .Host }}:443/images/logo.gif" kind="splash"/>
   <shortcut online="true"/>
 </information>
 <application-desc main-class="com.avocent.idrac.kvm.Main">
   <argument>ip={{ .Host }}</argument>
   <argument>vm=1</argument>
   <argument>helpurl=https://{{ .Host }}:443/help/contents.html</argument>
   <argument>title=INGG: {{ .Host }}</argument>
   <argument>user={{ .Username }}</argument>
   <argument>passwd={{ .Password }}</argument>
   <argument>kmport=5900</argument>
   <argument>vport=5900</argument>
   <argument>apcp=1</argument>
   <argument>chat=1</argument>
   <argument>F1=1</argument>
   <argument>custom=1</argument>
   <argument>scaling=15</argument>
   <argument>minwinheight=100</argument>
   <argument>minwinwidth=100</argument>
   <argument>videoborder=0</argument>
   <argument>version=2</argument>
 </application-desc>
 <security>
   <all-permissions/>
 </security>
 <resources>
   <j2se version="1.6+"/>
   <jar href="https://{{ .Host }}:443/software/avctKVM.jar" download="eager" main="true" />
 </resources>
 <resources os="Windows" arch="x86">
   <nativelib href="https://{{ .Host }}:443/software/avctKVMIOWin32.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLWin32.jar" download="eager"/>
 </resources>
 <resources os="Windows" arch="amd64">
   <nativelib href="https://{{ .Host }}:443/software/avctKVMIOWin64.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLWin64.jar" download="eager"/>
 </resources>
 <resources os="Windows" arch="x86_64">
   <nativelib href="https://{{ .Host }}:443/software/avctKVMIOWin64.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLWin64.jar" download="eager"/>
 </resources>
  <resources os="Linux" arch="x86">
    <nativelib href="https://{{ .Host }}:443/software/avctKVMIOLinux32.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLLinux32.jar" download="eager"/>
  </resources>
  <resources os="Linux" arch="i386">
    <nativelib href="https://{{ .Host }}:443/software/avctKVMIOLinux32.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLLinux32.jar" download="eager"/>
  </resources>
  <resources os="Linux" arch="i586">
    <nativelib href="https://{{ .Host }}:443/software/avctKVMIOLinux32.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLLinux32.jar" download="eager"/>
  </resources>
  <resources os="Linux" arch="i686">
    <nativelib href="https://{{ .Host }}:443/software/avctKVMIOLinux32.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLLinux32.jar" download="eager"/>
  </resources>
  <resources os="Linux" arch="amd64">
    <nativelib href="https://{{ .Host }}:443/software/avctKVMIOLinux64.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLLinux64.jar" download="eager"/>
  </resources>
  <resources os="Linux" arch="x86_64">
    <nativelib href="https://{{ .Host }}:443/software/avctKVMIOLinux64.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLLinux64.jar" download="eager"/>
  </resources>
  <resources os="Mac OS X" arch="x86_64">
    <nativelib href="https://{{ .Host }}:443/software/avctKVMIOMac64.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMAPI_DLLMac64.jar" download="eager"/>
  </resources>
</jnlp>
`
