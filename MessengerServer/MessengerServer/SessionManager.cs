using System.Text;
using RabbitMQ.Client.Events;

namespace MessengerServer
{
    class SessionManager
    {
        RabbitMQController _rmqCtrl;
        CommandExecutor _cmdExec;
        Config _config;

        public SessionManager(Config config)
        {
            _config = config;
            _cmdExec = new CommandExecutor(config.DbServer,
                config.DbUserName, config.DbUserPass, config.DbName);

            _rmqCtrl = new RabbitMQController(config.RMQUserName,
                config.RMQUserPass, config.VirtualHost, config.HostName);
        }

        public void Start()
        {
            SettingsRMQController();
            _rmqCtrl.GetMsg += ExecuteAndSendResult;
            _rmqCtrl.Listen(_config.InputQueue.Name);
        }

        void SettingsRMQController()
        {
            _rmqCtrl.GetChannel(_config.InputQueue.Exchange,
                _config.InputQueue.Name, _config.InputQueue.RoatKey);
            _rmqCtrl.GetChannel(_config.OutputQueue.Exchange,
                _config.OutputQueue.Name, _config.OutputQueue.RoatKey);
        }

        void ExecuteAndSendResult(object sender, BasicDeliverEventArgs e)
        {
            string queryResult = _cmdExec.SwitchAndExecuteCommand(Encoding.UTF8.GetString(e.Body));
            _rmqCtrl.SendMessage(queryResult, e.Exchange, _config.OutputQueue.Name);
        }
    }
}
