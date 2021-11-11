package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// VirtualPhysicalConnection is a nested struct in vpc response
type VirtualPhysicalConnection struct {
	Type                            string `json:"Type" xml:"Type"`
	Status                          string `json:"Status" xml:"Status"`
	CreationTime                    string `json:"CreationTime" xml:"CreationTime"`
	AdLocation                      string `json:"AdLocation" xml:"AdLocation"`
	PortNumber                      string `json:"PortNumber" xml:"PortNumber"`
	Spec                            string `json:"Spec" xml:"Spec"`
	ChargeType                      string `json:"ChargeType" xml:"ChargeType"`
	Description                     string `json:"Description" xml:"Description"`
	Bandwidth                       int64  `json:"Bandwidth" xml:"Bandwidth"`
	EnabledTime                     string `json:"EnabledTime" xml:"EnabledTime"`
	LineOperator                    string `json:"LineOperator" xml:"LineOperator"`
	PeerLocation                    string `json:"PeerLocation" xml:"PeerLocation"`
	RedundantPhysicalConnectionId   string `json:"RedundantPhysicalConnectionId" xml:"RedundantPhysicalConnectionId"`
	Name                            string `json:"Name" xml:"Name"`
	CircuitCode                     string `json:"CircuitCode" xml:"CircuitCode"`
	EndTime                         string `json:"EndTime" xml:"EndTime"`
	PortType                        string `json:"PortType" xml:"PortType"`
	BusinessStatus                  string `json:"BusinessStatus" xml:"BusinessStatus"`
	LoaStatus                       string `json:"LoaStatus" xml:"LoaStatus"`
	AccessPointId                   string `json:"AccessPointId" xml:"AccessPointId"`
	PhysicalConnectionId            string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
	ProductType                     string `json:"ProductType" xml:"ProductType"`
	ParentPhysicalConnectionId      string `json:"ParentPhysicalConnectionId" xml:"ParentPhysicalConnectionId"`
	VirtualPhysicalConnectionStatus string `json:"VirtualPhysicalConnectionStatus" xml:"VirtualPhysicalConnectionStatus"`
	ParentPhysicalConnectionAliUid  string `json:"ParentPhysicalConnectionAliUid" xml:"ParentPhysicalConnectionAliUid"`
	OrderMode                       string `json:"OrderMode" xml:"OrderMode"`
	AliUid                          string `json:"AliUid" xml:"AliUid"`
	VlanId                          string `json:"VlanId" xml:"VlanId"`
}