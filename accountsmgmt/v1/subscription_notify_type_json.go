/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalSubscriptionNotify writes a value of the 'subscription_notify' type to the given writer.
func MarshalSubscriptionNotify(object *SubscriptionNotify, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeSubscriptionNotify(object, stream)
	stream.Flush()
	return stream.Error
}

// writeSubscriptionNotify writes a value of the 'subscription_notify' type to the given stream.
func writeSubscriptionNotify(object *SubscriptionNotify, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if object.bccAddress != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("bcc_address")
		stream.WriteString(*object.bccAddress)
		count++
	}
	if object.subject != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("subject")
		stream.WriteString(*object.subject)
		count++
	}
	if object.templateName != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("template_name")
		stream.WriteString(*object.templateName)
		count++
	}
	if object.templateParameters != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("template_parameters")
		writeTemplateParameterList(object.templateParameters, stream)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalSubscriptionNotify reads a value of the 'subscription_notify' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalSubscriptionNotify(source interface{}) (object *SubscriptionNotify, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readSubscriptionNotify(iterator)
	err = iterator.Error
	return
}

// readSubscriptionNotify reads a value of the 'subscription_notify' type from the given iterator.
func readSubscriptionNotify(iterator *jsoniter.Iterator) *SubscriptionNotify {
	object := &SubscriptionNotify{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "bcc_address":
			value := iterator.ReadString()
			object.bccAddress = &value
		case "subject":
			value := iterator.ReadString()
			object.subject = &value
		case "template_name":
			value := iterator.ReadString()
			object.templateName = &value
		case "template_parameters":
			value := readTemplateParameterList(iterator)
			object.templateParameters = value
		default:
			iterator.ReadAny()
		}
	}
	return object
}