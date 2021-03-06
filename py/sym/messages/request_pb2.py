# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: sym/messages/request.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from sym.models import resource_pb2 as sym_dot_models_dot_resource__pb2
from sym.models import schema_pb2 as sym_dot_models_dot_schema__pb2
from sym.models import user_pb2 as sym_dot_models_dot_user__pb2
from sym.models import uuid_pb2 as sym_dot_models_dot_uuid__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='sym/messages/request.proto',
  package='sym.messages',
  syntax='proto3',
  serialized_options=b'Z%github.com/symopsio/types/go/messages',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1asym/messages/request.proto\x12\x0csym.messages\x1a\x19sym/models/resource.proto\x1a\x17sym/models/schema.proto\x1a\x15sym/models/user.proto\x1a\x15sym/models/uuid.proto\"\x8d\x02\n\x07Request\x12\x1c\n\x02id\x18\x01 \x01(\x0b\x32\x10.sym.models.UUID\x12\"\n\x06schema\x18\x02 \x01(\x0b\x32\x12.sym.models.Schema\x12,\n\x06target\x18\x03 \x01(\x0b\x32\x1c.sym.messages.Request.Target\x12(\n\x04meta\x18\x04 \x01(\x0b\x32\x1a.sym.messages.Request.Meta\x1a\x16\n\x04Meta\x12\x0e\n\x06reason\x18\x01 \x01(\t\x1aP\n\x06Target\x12\x1e\n\x04user\x18\x01 \x01(\x0b\x32\x10.sym.models.User\x12&\n\x08resource\x18\x02 \x01(\x0b\x32\x14.sym.models.ResourceB\'Z%github.com/symopsio/types/go/messagesb\x06proto3'
  ,
  dependencies=[sym_dot_models_dot_resource__pb2.DESCRIPTOR,sym_dot_models_dot_schema__pb2.DESCRIPTOR,sym_dot_models_dot_user__pb2.DESCRIPTOR,sym_dot_models_dot_uuid__pb2.DESCRIPTOR,])




_REQUEST_META = _descriptor.Descriptor(
  name='Meta',
  full_name='sym.messages.Request.Meta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='reason', full_name='sym.messages.Request.Meta.reason', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=308,
  serialized_end=330,
)

_REQUEST_TARGET = _descriptor.Descriptor(
  name='Target',
  full_name='sym.messages.Request.Target',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='sym.messages.Request.Target.user', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource', full_name='sym.messages.Request.Target.resource', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=332,
  serialized_end=412,
)

_REQUEST = _descriptor.Descriptor(
  name='Request',
  full_name='sym.messages.Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='sym.messages.Request.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='schema', full_name='sym.messages.Request.schema', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='target', full_name='sym.messages.Request.target', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='meta', full_name='sym.messages.Request.meta', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_REQUEST_META, _REQUEST_TARGET, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=143,
  serialized_end=412,
)

_REQUEST_META.containing_type = _REQUEST
_REQUEST_TARGET.fields_by_name['user'].message_type = sym_dot_models_dot_user__pb2._USER
_REQUEST_TARGET.fields_by_name['resource'].message_type = sym_dot_models_dot_resource__pb2._RESOURCE
_REQUEST_TARGET.containing_type = _REQUEST
_REQUEST.fields_by_name['id'].message_type = sym_dot_models_dot_uuid__pb2._UUID
_REQUEST.fields_by_name['schema'].message_type = sym_dot_models_dot_schema__pb2._SCHEMA
_REQUEST.fields_by_name['target'].message_type = _REQUEST_TARGET
_REQUEST.fields_by_name['meta'].message_type = _REQUEST_META
DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Request = _reflection.GeneratedProtocolMessageType('Request', (_message.Message,), {

  'Meta' : _reflection.GeneratedProtocolMessageType('Meta', (_message.Message,), {
    'DESCRIPTOR' : _REQUEST_META,
    '__module__' : 'sym.messages.request_pb2'
    # @@protoc_insertion_point(class_scope:sym.messages.Request.Meta)
    })
  ,

  'Target' : _reflection.GeneratedProtocolMessageType('Target', (_message.Message,), {
    'DESCRIPTOR' : _REQUEST_TARGET,
    '__module__' : 'sym.messages.request_pb2'
    # @@protoc_insertion_point(class_scope:sym.messages.Request.Target)
    })
  ,
  'DESCRIPTOR' : _REQUEST,
  '__module__' : 'sym.messages.request_pb2'
  # @@protoc_insertion_point(class_scope:sym.messages.Request)
  })
_sym_db.RegisterMessage(Request)
_sym_db.RegisterMessage(Request.Meta)
_sym_db.RegisterMessage(Request.Target)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
