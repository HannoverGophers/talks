FUNC=golambdafunctionimage

for i in {1..5}; do
  # 1) Dummy-Env ändern -> erzwingt neue Umgebungen
  aws lambda update-function-configuration \
    --function-name "$FUNC" \
    --environment "Variables={DUMMY_TS=$(date +%s%N)}" >/dev/null

  # 2) Auf „LastUpdateStatus=Successful“ warten (wichtig!)
  until [[ "$(aws lambda get-function-configuration \
        --function-name "$FUNC" \
        --query 'LastUpdateStatus' --output text)" == "Successful" ]]; do
    sleep 1
  done

  # 3) Invoke + REPORT ziehen
  base64 --decode <<<"$(aws lambda invoke \
      --function-name "$FUNC" \
      --payload '{}' \
      --log-type Tail \
      /dev/null \
      --query 'LogResult' --output text)" \
    | grep -E "Duration: .*Billed Duration" \
    >> benchmark_cold.log
done