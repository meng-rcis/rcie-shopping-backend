import fs from "fs";
import csv from "csv-parser";
// @ts-ignore
import { parse } from "json2csv";
import { findNearestFailedTime } from "../utils";

export const streamGetFailedResponse = async (log: string): Promise<any[]> => {
  let countLog = 0;
  const failedResponse: any[] = [];

  return new Promise(function (resolve, reject) {
    fs.createReadStream(log)
      .pipe(
        csv({
          separator: ",",
          headers: [
            "timeStamp",
            "elapsed",
            "label",
            "responseCode",
            "responseMessage",
            "threadName",
            "dataType",
            "success",
            "failureMessage",
            "bytes",
            "sentBytes",
            "grpThreads",
            "allThreads",
            "URL",
            "Latency",
            "IdleTime",
            "Connect",
          ],
        })
      )
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
  errorResponse: any[]
): Promise<void> => {
  let count = 0;
  const collection: any[] = [];
  const sortedResponse = errorResponse.sort(
    (a, b) => Number(a.timeStamp) - Number(b.timeStamp)
  );

  return new Promise(function (resolve, reject) {
    fs.createReadStream(currentFile)
      .pipe(
        csv({
          separator: ",",
          headers: [
            "Time",
            "iowait",
            "irq",
            "nice",
            "softirq",
            "steal",
            "system",
            "user",
            "Used",
            "br-10cd1f6f3899 receive",
            "docker0 receive",
            "eth0 receive",
            "eth1 receive",
            "lo receive",
            "veth6750e95 receive",
            "timeStamp",
          ],
        })
      )
      .on("data", (data: any) => {
        if (count === 0) {
          count++;
          return;
        }
        data.timeToFail = findNearestFailedTime(data.timeStamp, sortedResponse);
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