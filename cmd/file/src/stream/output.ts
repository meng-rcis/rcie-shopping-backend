import fs from "fs";
import csv from "csv-parser";
// @ts-ignore
import { parse } from "json2csv";
import { findNearestFailedTime } from "../utils/timestamp";
import {
  TIMESTAMP_HEADER,
  TIME_HEADER,
  TIME_TO_FAIL_HEADER,
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
  const metricsHeaders = metrics.flatMap((metric) => metric.headers);
  const sortedErrors = errorResponse.sort((a, b) => a.timeStamp - b.timeStamp);

  return new Promise(function (resolve, reject) {
    fs.createReadStream(currentFile)
      .pipe(
        csv({
          separator: ",",
          headers: [TIME_HEADER, TIMESTAMP_HEADER, ...metricsHeaders],
        })
      )
      .on("data", (data: any) => {
        if (count === 0) {
          count++;
          return;
        }
        data[TIME_TO_FAIL_HEADER] = findNearestFailedTime(
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
