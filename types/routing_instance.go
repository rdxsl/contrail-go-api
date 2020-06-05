//
// Automatically generated. DO NOT EDIT.
//

package types

import (
	"encoding/json"

	"github.com/Juniper/contrail-go-api"
)

const (
	routing_instance_id_perms uint64 = 1 << iota
	routing_instance_display_name
	routing_instance_virtual_machine_interface_back_refs
	routing_instance_virtual_network_back_refs
	routing_instance_route_target_refs
)

type RoutingInstance struct {
	contrail.ObjectBase
	id_perms                            IdPermsType
	display_name                        string
	virtual_machine_interface_back_refs contrail.ReferenceList
	virtual_network_back_refs           contrail.ReferenceList
	route_target_refs                   contrail.ReferenceList
	valid                               uint64
	modified                            uint64
	originalMap                         map[string]contrail.ReferenceList
}

func (obj *RoutingInstance) GetType() string {
	return "routing-instance"
}

func (obj *RoutingInstance) GetDefaultParent() []string {
	name := []string{"default-domain", "default-project", "default-virtual-network"}
	return name
}

func (obj *RoutingInstance) GetDefaultParentType() string {
	return "virtual-network"
}

func (obj *RoutingInstance) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *RoutingInstance) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *RoutingInstance) addChange(
	name string, refList contrail.ReferenceList) {
	if obj.originalMap == nil {
		obj.originalMap = make(map[string]contrail.ReferenceList)
	}
	var refCopy contrail.ReferenceList
	copy(refCopy, refList)
	obj.originalMap[name] = refCopy
}

func (obj *RoutingInstance) UpdateDone() {
	obj.modified = 0
	obj.originalMap = nil
}

func (obj *RoutingInstance) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *RoutingInstance) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified |= routing_instance_id_perms
}

func (obj *RoutingInstance) GetDisplayName() string {
	return obj.display_name
}

func (obj *RoutingInstance) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified |= routing_instance_display_name
}

func (obj *RoutingInstance) readVirtualMachineInterfaceBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid&routing_instance_virtual_machine_interface_back_refs == 0) {
		err := obj.GetField(obj, "virtual_machine_interface_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *RoutingInstance) GetVirtualMachineInterfaceBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaceBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interface_back_refs, nil
}

func (obj *RoutingInstance) readVirtualNetworkBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid&routing_instance_virtual_network_back_refs == 0) {
		err := obj.GetField(obj, "virtual_network_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *RoutingInstance) GetVirtualNetworkBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualNetworkBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_network_back_refs, nil
}

func (obj *RoutingInstance) readRouteTargetRefs() error {
	if !obj.IsTransient() &&
		(obj.valid&routing_instance_route_target_refs == 0) {
		err := obj.GetField(obj, "route_target_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *RoutingInstance) GetRouteTargetRefs() (contrail.ReferenceList, error) {
	err := obj.readRouteTargetRefs()
	if err != nil {
		return nil, err
	}
	return obj.route_target_refs, nil
}

func (obj *RoutingInstance) SetRouteTargetRefs(value contrail.ReferenceList) {
	obj.route_target_refs = value
	obj.modified |= routing_instance_route_target_refs
}

func (obj *RoutingInstance) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified&routing_instance_id_perms != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified&routing_instance_display_name != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}
	if obj.modified&routing_instance_route_target_refs != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.route_target_refs)
		if err != nil {
			return nil, err
		}
		msg["route_target_refs"] = &value
	}

	return json.Marshal(msg)
}

func (obj *RoutingInstance) UnmarshalJSON(body []byte) error {
	var m map[string]json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	err = obj.UnmarshalCommon(m)
	if err != nil {
		return err
	}
	for key, value := range m {
		switch key {
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid |= routing_instance_id_perms
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid |= routing_instance_display_name
			}
			break
		case "virtual_network_back_refs":
			err = json.Unmarshal(value, &obj.virtual_network_back_refs)
			if err == nil {
				obj.valid |= routing_instance_virtual_network_back_refs
			}
			break
		case "virtual_machine_interface_back_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr PolicyBasedForwardingRuleType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid |= routing_instance_virtual_machine_interface_back_refs
				obj.virtual_machine_interface_back_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.virtual_machine_interface_back_refs = append(obj.virtual_machine_interface_back_refs, ref)
				}
				break
			}
		case "route_target_refs":
			type ReferenceElement struct {
				To   []string
				Uuid string
				Href string
				Attr RouteTargetRefAttr
			}
			var array []ReferenceElement
			err = json.Unmarshal(value, &array)
			if err != nil {
				break
			}
			obj.valid |= routing_instance_route_target_refs
			obj.route_target_refs = make(contrail.ReferenceList, 0)
			for _, element := range array {
				ref := contrail.Reference{
					element.To,
					element.Uuid,
					element.Href,
					element.Attr,
				}
				obj.route_target_refs = append(obj.route_target_refs, ref)
			}
			break
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *RoutingInstance) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified&routing_instance_id_perms != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified&routing_instance_display_name != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}
	if obj.modified&routing_instance_route_target_refs != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.route_target_refs)
		if err != nil {
			return nil, err
		}
		msg["route_target_refs"] = &value
	}
	return json.Marshal(msg)
}

func (obj *RoutingInstance) UpdateReferences() error {

	return nil
}

func RoutingInstanceByName(c contrail.ApiClient, fqn string) (*RoutingInstance, error) {
	obj, err := c.FindByName("routing-instance", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*RoutingInstance), nil
}

func RoutingInstanceByUuid(c contrail.ApiClient, uuid string) (*RoutingInstance, error) {
	obj, err := c.FindByUuid("routing-instance", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*RoutingInstance), nil
}
