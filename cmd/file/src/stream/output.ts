import fs from "fs";
import csv from "csv-parser";
// @ts-ignore
import { parse } from "json2csv";
import { findNearestFailedTime } from "../utils";
import {
  CPU_HEADER,
  MEM_HEADER,
  NET_HEADER,
  TIMESTAMP_HEADER,
  TIME_HEADER,
} from "../constant/metrics-header";
import { LOG_HEADER } from "../constant/log-header";
import { IMetric } from "../interfaces/metric";

export const streamGetFailedResponse = async (log: string): Promise<any[]> => {
  let countLog = 0;
  const failedResponse: any[] = [];

  return new Promise(function (resolve, reject) {
    fs.createReadStream(log)
      .pipe(csv({ separator: ",", headers: LOG_HEADER }))
      .on("data", (data: any) => {
        if (countLog === 0) {
          countLog++;
          return;
        }
        if (data.responseCode !== "200") {
          failedResponse.push(data);
        }
      })
      .on("end", () => {
        resolve(failedResponse);
      })
      .on("error", reject);
  });
};

export const streamAddTimeToFail = async (
  currentFile: string,
  newFile: string,
  errorResponse: any[],
  metrics: IMetric[]
): Promise<void> => {
  let count = 0;
  const collection: any[] = [];
  const metricsHeader = metrics.flatMap((metric) => metric.headers);
  const sortedErrors = errorResponse.sort((a, b) => a.timeStamp - b.timeStamp);

  return new Promise(function (resolve, reject) {
    fs.createReadStream(currentFile)
      .pipe(
        csv({
          separator: ",",
          headers: [TIME_HEADER, TIMESTAMP_HEADER, ...metricsHeader],
        })
      )
      .on("data", (data: any) => {
        if (count === 0) {
          count++;
          return;
        }
        data.timeToFail = findNearestFailedTime(
          data[TIMESTAMP_HEADER],
          sortedErrors
        );
        collection.push(data);
      })
      .on("end", () => {
        var result = parse(collection);
        fs.writeFileSync(newFile, result);
        resolve();
      })
      .on("error", reject);
  });
};
