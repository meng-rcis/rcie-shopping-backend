import { IPath } from "./interfaces/path";
import {
  streamAddMetricsTimeStamp,
  streamAddTimeToFail,
  streamGetFailedResponse,
} from "./stream";

export const addMetricsTimeStamp = async (metrics: string): Promise<string> => {
  const newFile = metrics.replace(".csv", "_with_output.csv");
  await streamAddMetricsTimeStamp(metrics, newFile);
  return newFile;
};

export const addTimeToFail = async (
  log: string,
  metrics: string
): Promise<string> => {
  const errorResponse: any[] = await streamGetFailedResponse(log);
  await streamAddTimeToFail(metrics, metrics, errorResponse);
  return metrics;
};

export const addTimestamp = async (path: IPath): Promise<string> => {
  const { logFile, metricsFile } = path;
  const metricsWithTimeStamp = await addMetricsTimeStamp(metricsFile);
  const metricsWithTimeToFail = await addTimeToFail(
    logFile,
    metricsWithTimeStamp
  );

  return metricsWithTimeToFail;
};
