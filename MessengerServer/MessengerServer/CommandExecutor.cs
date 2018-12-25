using System.Collections.Generic;
using MessengerServer.DbWorkers;

namespace MessengerServer
{
    class CommandExecutor
    {
        readonly IDbWorker _dbWorker;

        public CommandExecutor(string dbServer, string userName, string userPass, string dbName) :
            this(new PgsqlDbWorker(dbServer, userName, userPass, dbName))
        {
        }

        public CommandExecutor(IDbWorker dbWorker)
        {
            _dbWorker = dbWorker;
        }

        public string SwitchAndExecuteCommand(string json)
        {
            return SwitchAndExecuteCommand(JsonParser.Deserialize<Request>(json));
        }

        public string SwitchAndExecuteCommand(Request request)
        {
            string result;

            switch (request.Command)
            {
                case "enter_dialog":
                    result = EnterDialog(request.Data);
                    break;
                case "create_dialog":
                    result = CreateDialog(request.Data);
                    break;
                case "leave_dialog":
                    result = LeaveDialog(request.Data);
                    break;
                case "invite_to_dialog":
                    result = InviteUserToDialog(request.Data);
                    break;
                case "get_dialogs":
                    result = GetDialogsList(request.Data);
                    break;
                case "auth":
                    result = Authentification(request.Data);
                    break;
                case "reg":
                    result = Registration(request.Data);
                    break;
                default:
                    result = JsonParser.SerializeException("Такой команды не существует");
                    break;
            }

            request.Data = result;
            return JsonParser.Serialize(request);
        }

        public string EnterDialog(string obj)
        {
            var data = JsonParser.DeserializeAnonymousType(obj,
                new { DialogId = 0L, UserID = 0L });
            string res = "";

            bool exist = _dbWorker.UserInDialog(data.UserID, data.DialogId);
            if (!exist)
                res = "Вы не являетесь участником диалога.";

            return JsonParser.SerializeException(res);
        }

        public string CreateDialog(string obj)
        {
            var data = JsonParser.DeserializeAnonymousType(obj,
                new { DialogName = "", FromUserId = 0L, UserLogins = new string[0] });
            var userIds = new List<long>() { data.FromUserId };
            var notExistNames = new List<string>();
            string exc = "";

            if (data.UserLogins != null)
                foreach (var item in data.UserLogins)
                    if (_dbWorker.GetUserId(item, out long id))
                        userIds.Add(id);
                    else
                        notExistNames.Add(item);

            long dialogId = _dbWorker.CreateDialog(data.DialogName, userIds);
            if (notExistNames.Count != 0)
                exc = string.Join(", ", notExistNames);

            return JsonParser.Serialize(
                new { UserIds = userIds, DialogId = dialogId, Exception = exc });
        }

        public string LeaveDialog(string obj)
        {
            var data = JsonParser.DeserializeAnonymousType(obj,
                new { DialogId = 0L, UserID = 0L });

            _dbWorker.DeleteUserFromDialog(data.UserID, data.DialogId);

            return JsonParser.SerializeException("");
        }

        public string InviteUserToDialog(string obj)
        {
            var data = JsonParser.DeserializeAnonymousType(obj,
                new { DialogId = 0L, FromUserID = 0L, ToUserLogin = "" });
            string res = "";

            // Проверить есть ли такой FromUserID в списке DialogId
            if (!_dbWorker.UserInDialog(data.FromUserID, data.DialogId))
                res = "Такого диалога не существует или вы не являетесь его участником.";

            // Получить id user`а по логину// может не быть
            if (!_dbWorker.GetUserId(data.ToUserLogin, out long toUserId))
                res = "Пользователя с таким логином не существует.";

            // Проверить есть ли такой toUserId в списке DialogId
            if (!_dbWorker.UserInDialog(toUserId, data.DialogId))
                _dbWorker.AddUserInDialog(toUserId, data.DialogId);

            return JsonParser.Serialize(new { UserId = toUserId, Exception = res });
        }

        public string GetDialogsList(string obj)
        {
            var data = JsonParser.DeserializeAnonymousType(obj,
                new { UserID = 0L });
            var dialogs = _dbWorker.GetAllUserDialogs(data.UserID);

            return JsonParser.Serialize(new { DialogsList = dialogs, Exception = "" });
        }

        public string Authentification(string obj)
        {
            var data = JsonParser.DeserializeAnonymousType(obj,
                new { UserId = 0L, Password = "" });
            string res = "";

            bool auth = _dbWorker.Auth(data.UserId, data.Password);
            if (!auth)
                res = "Пользователя с таким id и паролем не существует.";

            return JsonParser.SerializeException(res);
        }

        public string Registration(string obj)
        {
            var data = JsonParser.DeserializeAnonymousType(obj,
                new { UserId = 0L, Login = "", Password = "" });
            string res = "";

            bool exist = _dbWorker.UserIsExist(data.UserId);
            if (exist)
                res = "Пользователь с таким id уже зарегистрирован.";
            else
                _dbWorker.CreateUser(data.UserId, data.Login, data.Password);

            return JsonParser.SerializeException(res);
        }
    }
}
