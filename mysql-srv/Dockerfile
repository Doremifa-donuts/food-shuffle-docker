FROM mysql:latest

# 環境変数を使用してMySQLのポートを指定
ARG MYSQL_PORT=3308

# my.cnfにポートを追記または追加する処理
RUN bash -c ' \
    MY_CNF_PATH="/etc/my.cnf"; \
    if grep -q "^\[mysqld\]" "$MY_CNF_PATH"; then \
        # [mysqld]セクションが存在する場合、port設定を削除して追加
        sed -i "/^\[mysqld\]/,/^\[/{/port/d}" "$MY_CNF_PATH" && \
        sed -i "/^\[mysqld\]/a port = ${MYSQL_PORT}" "$MY_CNF_PATH"; \
    else \
        # [mysqld]セクションが存在しない場合、新たにセクションを追加
        echo -e "\n[mysqld]\nport = ${MYSQL_PORT}" >> "$MY_CNF_PATH"; \
    fi \
'

# MySQLデーモンを起動
CMD ["mysqld"]
