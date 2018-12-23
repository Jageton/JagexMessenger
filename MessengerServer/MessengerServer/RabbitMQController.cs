using System;
using System.Text;
using RabbitMQ.Client;
using RabbitMQ.Client.MessagePatterns;
using RabbitMQ.Client.Events;

namespace MessengerServer
{
    class RabbitMQController
    {
        public delegate void RMQGetMsg(object sender, BasicDeliverEventArgs e);

        public event RMQGetMsg GetMsg;

        readonly IConnection _connection;
        readonly IModel _model;

        public RabbitMQController(string userName, string password, string virtualHost, string hostName)
        {
            _connection = GetConnection(userName, password, virtualHost, hostName);
            _model = _connection.CreateModel();
        }

        IConnection GetConnection(string userName, string password, string virtualHost, string hostName)
        {
            ConnectionFactory factory = new ConnectionFactory
            {
                UserName = userName,
                Password = password,
                VirtualHost = virtualHost,
                HostName = hostName
            };
            IConnection conn = factory.CreateConnection();
            return conn;
        }
        
        public void SendMessage(string msg, string exchangeName, string routingKey)
        {
            SendMessage(Encoding.UTF8.GetBytes(msg), exchangeName, routingKey);
        }

        public void SendMessage(byte[] msg, string exchangeName, string routingKey)
        {
            _model.BasicPublish(exchangeName, routingKey, null, msg);
        }

        public void GetChannel(string exchangeName, string queueName, string routingKey)
        {
            _model.ExchangeDeclare(exchangeName, ExchangeType.Direct);
            _model.QueueDeclare(queueName, false, false, false, null);
            _model.QueueBind(queueName, exchangeName, routingKey, null);
        }

        public void Listen(string queueName)
        {
            var subscription = new Subscription(_model, queueName, false);

            while (true)
            {
                Console.WriteLine("Listen");
                    BasicDeliverEventArgs eventArgs = subscription.Next();
                OnGetMsg(eventArgs);
                subscription.Ack(eventArgs);
            }
        }

        protected virtual void OnGetMsg(BasicDeliverEventArgs e)
        {
            GetMsg?.Invoke(this, e);
        }
    }
}
