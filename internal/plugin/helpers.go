package plugin

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func scopedDescriptorTypeName(pkg protoreflect.FullName, desc protoreflect.Descriptor) string {
	name := string(desc.Name())
	var prefix string
	if desc.Parent() != desc.ParentFile() {
		prefix = descriptorTypeName(desc.Parent())
	}
	if desc.ParentFile().Package() != pkg {
		prefix = packagePrefix(desc.ParentFile().Package()) + prefix
	}
	return prefix + name
}

func descriptorTypeName(desc protoreflect.Descriptor) string {
	name := string(desc.Name())
	var prefix string
	if desc.Parent() != desc.ParentFile() {
		prefix = descriptorTypeName(desc.Parent())
	}
	return prefix + name
}

func packagePrefix(pkg protoreflect.FullName) string {
	out := ""
	for _, s := range strings.Split(string(pkg), ".") {
		s = strings.ToLower(s)
		if len(s) > 1 {
			out = out + strings.ToUpper(s[:1]) + s[1:]
			continue
		}
		out = out + strings.ToUpper(s)
	}
	return out
}

func rangeFields(message protoreflect.MessageDescriptor, f func(field protoreflect.FieldDescriptor)) {
	for i := 0; i < message.Fields().Len(); i++ {
		f(message.Fields().Get(i))
	}
}

func rangeMethods(methods protoreflect.MethodDescriptors, f func(method protoreflect.MethodDescriptor)) {
	for i := 0; i < methods.Len(); i++ {
		f(methods.Get(i))
	}
}

func rangeEnumValues(enum protoreflect.EnumDescriptor, f func(value protoreflect.EnumValueDescriptor, last bool)) {
	for i := 0; i < enum.Values().Len(); i++ {
		if i == enum.Values().Len()-1 {
			f(enum.Values().Get(i), true)
		} else {
			f(enum.Values().Get(i), false)
		}
	}
}

func t(n int) string {
	return strings.Repeat("  ", n)
}
