package main

const viewer6 string = `
<?xml version="1.0" encoding="UTF-8"?>
<jnlp codebase="https://{{ .Host }}:443" spec="1.0+">
<information>
  <title>iDRAC6 Console Redirection Client</title>
  <vendor>Dell Inc.</vendor>
   <icon href="https://{{ .Host }}:443/images/logo.gif" kind="splash"/>
   <shortcut online="true"/>
 </information>
 <application-desc main-class="com.avocent.idrac.kvm.Main">
   <argument>title=DRAC KVM: {{ .Host }}</argument>
   <argument>ip={{ .Host }}</argument>
   <argument>vmprivilege=true</argument>
   <argument>helpurl=https://{{ .Host }}:443/help/contents.html</argument>
   <argument>user={{ .Username }}</argument>
   <argument>passwd={{ .Password }}</argument>
   <argument>kmport=5900</argument>
   <argument>vport=5900</argument>
   <argument>apcp=1</argument>
   <argument>version=2</argument>
 </application-desc>
 <security>
   <all-permissions/>
 </security>
 <resources>
   <j2se version="1.6 1.5 1.4+"/>
   <jar href="https://{{ .Host }}:443/software/avctKVM.jar" download="eager" main="true" />
   <jar href="https://{{ .Host }}:443/software/jpcsc.jar" download="eager"/>
 </resources>
 <resources os="Windows">
   <nativelib href="https://{{ .Host }}:443/software/avctKVMIOWin32.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMWin32.jar" download="eager"/>
 </resources>
  <resources os="Linux">
    <nativelib href="https://{{ .Host }}:443/software/avctKVMIOLinux.jar" download="eager"/>
   <nativelib href="https://{{ .Host }}:443/software/avctVMLinux.jar" download="eager"/>
  </resources>
</jnlp>
`
