using System.Collections.Generic;
using System.Text;
using Npgsql;

namespace MessengerServer.DbWorkers
{
    class PgsqlDbWorker : IDbWorker
    {
        NpgsqlConnection _connect;

        public PgsqlDbWorker(string dbServer, string userName, string userPass, string dbName)
        {
            var connString = $"Server={ dbServer};User Id={ userName};" +
                $"Password={userPass};Database={dbName};Encoding=UTF8;";
            _connect = new NpgsqlConnection(connString);
        }

        /// <summary>
        /// Попытается открыть подключение к бд, если
        /// попытка успешна - закроет, иначе вызовет исключение.
        /// </summary>
        public void CheckConnection()
        {
            _connect.Open();
            _connect.Close();
        }

        public bool UserIsExist(long userId)
        {
            return CheckExist("SELECT messenger_sch.users_tb.id" +
                " FROM messenger_sch.users_tb " +
                $"WHERE messenger_sch.users_tb.id = {userId}");
        }

        public void CreateUser(long userId, string login, string password)
        {
            ExecuteNonQuery("INSERT INTO messenger_sch.users_tb(" + 
                "id, login, password, dialogs_id)" + 
                $"VALUES ({userId}, '{login}', '{password}', ARRAY[]::bigint[])");
        }

        public bool UserInDialog(long userId, long dialogId)
        {
            return CheckExist("SELECT messenger_sch.dialogs_tb.id" +
                " FROM messenger_sch.dialogs_tb " +
                $"WHERE messenger_sch.dialogs_tb.id = {dialogId} AND " +
                $"{userId} = ANY (messenger_sch.dialogs_tb.users_id)");
        }

        public void AddUserInDialog(long userId, long dialogId)
        {
            AddUserIdInDialog(userId, dialogId);
            AddDialogIdInUser(userId, dialogId);
        }

        void AddUserIdInDialog(long userId, long dialogId)
        {
            var ids = new List<long>(GetAllDialogUsersId(dialogId));
            if (ids == null)
                return;
            ids.Add(userId);
            AssertUsersId(dialogId, ids);
        }

        void AddDialogIdInUser(long userId, long dialogId)
        {
            var ids = new List<long>(GetAllUserDialogsId(userId));
            if (ids == null)
                return;
            ids.Add(dialogId);
            AssertDialogsId(userId, ids);
        }

        public long CreateDialog(string dialogName, ICollection<long> userIds)
        {
            string cmd = "INSERT INTO messenger_sch.dialogs_tb(name, users_id) " +
                $"VALUES ('{dialogName}', ARRAY[{string.Join(",", userIds)}])";
            ExecuteNonQuery(cmd);

            long dialogId = (long)ExecuteScalar(
                "SELECT id FROM messenger_sch.dialogs_tb ORDER BY id DESC LIMIT 1");
            
            foreach (var item in userIds)
                AddDialogIdInUser(item, dialogId);

            return dialogId;
        }

        public void DeleteUserFromDialog(long userId, long dialogId)
        {
            var ids = new List<long>(GetAllDialogUsersId(dialogId));
            if (ids == null)
                return;
            ids.Remove(userId);
            AssertUsersId(dialogId, ids);

            ids = new List<long>(GetAllUserDialogsId(userId));
            if (ids == null)
                return;
            ids.Remove(dialogId);
            AssertDialogsId(userId, ids);
        }

        void AssertUsersId(long dialogId, ICollection<long> ids)
        {
            ExecuteNonQuery("UPDATE messenger_sch.dialogs_tb " +
                $"SET users_id = ARRAY[{string.Join(",", ids)}] " +
                $"WHERE id = {dialogId}");
        }

        void AssertDialogsId(long userId, ICollection<long> ids)
        {
            ExecuteNonQuery("UPDATE messenger_sch.users_tb " +
                $"SET dialogs_id = ARRAY[{string.Join(",", ids)}] " +
                $"WHERE id = {userId}");
        }

        public long[] GetAllDialogUsersId(long dialogId)
        {
            return (long[])ExecuteScalar(
                "SELECT messenger_sch.dialogs_tb.users_id" +
                " FROM messenger_sch.dialogs_tb " +
                $"WHERE messenger_sch.dialogs_tb.id = {dialogId}");
        }

        public bool GetUserId(string login, out long userId)
        {
            userId = 0;
            var obj = ExecuteScalar("SELECT messenger_sch.users_tb.id" +
                $" FROM messenger_sch.users_tb WHERE messenger_sch.users_tb.login = '{login}'");
            if (obj == null)
                return false;
            userId = (long)obj;
            return true;
        }

        public ICollection<Dialog> GetAllUserDialogs(long userId)
        {
            long[] dialogsId = GetAllUserDialogsId(userId);
            if (dialogsId.Length == 0)
                return new List<Dialog>();
            
            string where = BuildQueryStringOfEquals("messenger_sch.dialogs_tb.id", dialogsId);
            string cmd = "SELECT messenger_sch.dialogs_tb.id, messenger_sch.dialogs_tb.name" +
                " FROM messenger_sch.dialogs_tb " +
                $"WHERE {where}";

            return GetDialogsInfo(cmd);
        }

        public long[] GetAllUserDialogsId(long userId)
        {
            return (long[])ExecuteScalar(
                "SELECT messenger_sch.users_tb.dialogs_id" +
                " FROM messenger_sch.users_tb " +
                $"WHERE messenger_sch.users_tb.id = {userId}");
        }

        string BuildQueryStringOfEquals<T>(string field, params T[] values)
        {
            var sb = new StringBuilder();
            sb.Append(field + " = " + values[0]);

            for (int i = 1; i < values.Length; i++)
                sb.Append(string.Format("OR {0} = {1}", field, values[i]));

            return sb.ToString();
        }

        public ICollection<Dialog> GetDialogsInfo(string command)
        {
            ICollection<Dialog> res = new List<Dialog>();
            _connect.Open();
            using (var cmd = new NpgsqlCommand(command, _connect))
            using (NpgsqlDataReader reader = cmd.ExecuteReader())
                while (reader.Read())
                    res.Add(new Dialog((long)reader.GetValue(0),
                        (string)reader.GetValue(1)));
            _connect.Close();
            return res;
        }

        public bool Auth(long userId, string password)
        {
            return CheckExist("SELECT messenger_sch.users_tb.id" +
                " FROM messenger_sch.users_tb " +
                $"WHERE messenger_sch.users_tb.id = {userId} AND " +
                $"messenger_sch.users_tb.password = '{password}'");
        }

        public bool CheckExist(string command)
        {
            bool check;
            _connect.Open();
            using (var cmd = new NpgsqlCommand(command, _connect))
            using (NpgsqlDataReader reader = cmd.ExecuteReader())
                check = reader.HasRows;
            _connect.Close();
            return check;
        }

        public object ExecuteScalar(string command)
        {
            object res;
            _connect.Open();
            using (var cmd = new NpgsqlCommand(command, _connect))
                res = cmd.ExecuteScalar();
            _connect.Close();
            return res;
        }

        public ICollection<T> ExecuteReader<T>(string command, int columnNumb = 0)
        {
            ICollection<T> res = new List<T>();
            _connect.Open();
            using (var cmd = new NpgsqlCommand(command, _connect))
            using (NpgsqlDataReader reader = cmd.ExecuteReader())
                while (reader.Read())
                    res.Add((T)reader.GetValue(columnNumb));
            _connect.Close();
            return res;
        }

        public int ExecuteNonQuery(string command)
        {
            int res;
            _connect.Open();
            using (var cmd = new NpgsqlCommand(command, _connect))
                res = cmd.ExecuteNonQuery();
            _connect.Close();
            return res;
        }
    }
}
