import fs from "fs";
import csv from "csv-parser";
// @ts-ignore
import { parse } from "json2csv";
import { IPath } from "./interface";
import { streamAddMetricsTimeStamp, streamGetFailedResponse } from "./stream";

export const addMetricsTimeStamp = async (metrics: string): Promise<string> => {
  const newFile = metrics.replace(".csv", "_formatted.csv");
  await streamAddMetricsTimeStamp(metrics, newFile);
  return newFile;
};

export const addTimeToFail = async (
  log: string,
  metrics: string
): Promise<string> => {
  const fileName = metrics.replace(".csv", "_with_log.csv");
  const failedResponse: any[] = await streamGetFailedResponse(log);
  const sortedResponse = failedResponse.sort(
    (a, b) => Number(a.timeStamp) - Number(b.timeStamp)
  );

  console.log("failedResponse", failedResponse);
  console.log("sortedResponse", sortedResponse);
  console.log("is same", failedResponse === sortedResponse);
  return fileName;
};

export const formatMetrics = async (path: IPath): Promise<string> => {
  const { logFile, metricsFile } = path;
  const metricsFileWithTimeStamp = await addMetricsTimeStamp(metricsFile);
  const result = await addTimeToFail(logFile, metricsFileWithTimeStamp);

  return "result";
};
