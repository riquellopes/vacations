# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from protons import vacations_pb2 as protons_dot_vacations__pb2


class VacationServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.WhenDoYouStartYourVacation = channel.unary_unary(
        '/protons.VacationService/WhenDoYouStartYourVacation',
        request_serializer=protons_dot_vacations__pb2.VacationRequest.SerializeToString,
        response_deserializer=protons_dot_vacations__pb2.VacationResponse.FromString,
        )


class VacationServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def WhenDoYouStartYourVacation(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_VacationServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'WhenDoYouStartYourVacation': grpc.unary_unary_rpc_method_handler(
          servicer.WhenDoYouStartYourVacation,
          request_deserializer=protons_dot_vacations__pb2.VacationRequest.FromString,
          response_serializer=protons_dot_vacations__pb2.VacationResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'protons.VacationService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
