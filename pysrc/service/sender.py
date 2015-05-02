import logging
import pika

from bunsan.broker import rabbit_pb2


class Sender(object):

    def __init__(self, channel, queue, identifier):
        self._logger = logging.getLogger(__name__)
        self._channel = channel
        self._queue = queue
        self._identifier = identifier
        self._properties = pika.BasicProperties(correlation_id=identifier)

    def send(self, body):
        self._logger.debug('Sending message to %s', self._queue)
        self._channel.basic_publish(exchange='',
                                    routing_key=self._queue,
                                    body=body,
                                    properties=self._properties)

    def sendmsg(self, format, *args):
        self.send((format % args).encode('utf8'))


class ProtoSender(Sender):

    def send_proto(self, proto):
        self.send(proto.SerializeToString())


class StatusSender(ProtoSender):

    def send_status(self, code, reason=None, data=None):
        status = rabbit_pb2.RabbitStatus()
        status.identifier = self._identifier
        status.status.code = code
        if reason is not None:
            status.status.reason = reason
        if data is not None:
            status.status.date = data
        self.send_proto(status)


class ResultSender(ProtoSender):

    def send_result(self, status, reason=None, data=None):
        result = rabbit_pb2.RabbitResult()
        result.identifier = self._identifier
        result.result.status = status
        if reason is not None:
            result.result.reason = reason
        if data is not None:
            result.result.data = data
        self.send_proto(result)


class ErrorSender(Sender):

    def __init__(self, channel, properties):
        super(ErrorSender, self).__init__(channel=channel,
                                          queue=properties.reply_to,
                                          identifier=properties.correlation_id)
