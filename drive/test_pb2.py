# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: test.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='test.proto',
  package='main',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\ntest.proto\x12\x04main\"X\n\x0b\x46ileRequest\x12\x10\n\x08\x66ilename\x18\x01 \x01(\t\x12\x11\n\ttimestamp\x18\x02 \x01(\x07\x12\x10\n\x08log_info\x18\x03 \x01(\t\x12\x12\n\x08trace_id\x18\x80\xf7\x02 \x01(\x05\x62\x06proto3')
)




_FILEREQUEST = _descriptor.Descriptor(
  name='FileRequest',
  full_name='main.FileRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='filename', full_name='main.FileRequest.filename', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='main.FileRequest.timestamp', index=1,
      number=2, type=7, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='log_info', full_name='main.FileRequest.log_info', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='trace_id', full_name='main.FileRequest.trace_id', index=3,
      number=48000, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=20,
  serialized_end=108,
)

DESCRIPTOR.message_types_by_name['FileRequest'] = _FILEREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FileRequest = _reflection.GeneratedProtocolMessageType('FileRequest', (_message.Message,), dict(
  DESCRIPTOR = _FILEREQUEST,
  __module__ = 'test_pb2'
  # @@protoc_insertion_point(class_scope:main.FileRequest)
  ))
_sym_db.RegisterMessage(FileRequest)


# @@protoc_insertion_point(module_scope)
