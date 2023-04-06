import { IPath } from "./interface";
import {
  streamAddMetricsTimeStamp,
  streamAddTimeToFail,
  streamGetFailedResponse,
} from "./stream";

export const addMetricsTimeStamp = async (metrics: string): Promise<string> => {
  const newFile = metrics.replace(".csv", "_formatted.csv");
  await streamAddMetricsTimeStamp(metrics, newFile);
  return newFile;
};

export const addTimeToFail = async (
  log: string,
  metrics: string
): Promise<string> => {
  const newFile = metrics.replace(".csv", "_with_log.csv");
  const failedResponse: any[] = await streamGetFailedResponse(log);
  const sortedResponse = failedResponse.sort(
    (a, b) => Number(a.timeStamp) - Number(b.timeStamp)
  );

  await streamAddTimeToFail(metrics, newFile, sortedResponse);
  return newFile;
};

export const formatMetrics = async (path: IPath): Promise<string> => {
  const { logFile, metricsFile } = path;
  const metricsWithTimeStamp = await addMetricsTimeStamp(metricsFile);
  const metricsWithLog = await addTimeToFail(logFile, metricsWithTimeStamp);

  return metricsWithLog;
};
