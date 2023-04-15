import fs from "fs";
import csv from "csv-parser";
// @ts-ignore
import { parse } from "json2csv";
import { IMetric } from "../interfaces/metric";
import { TIMESTAMP_HEADER, TIME_HEADER } from "../constant/metrics-header";

const streamGetBaseCollection = async (
  metric: IMetric,
  baseHeader: string
): Promise<any[]> => {
  let count = 0;
  const collection: any[] = [];

  return new Promise(function (resolve, reject) {
    fs.createReadStream(metric.path)
      .pipe(csv({ separator: ",", headers: [baseHeader, ...metric.headers] }))
      .on("data", (data: any) => {
        if (count === 0) {
          count++;
          return;
        }

        const base = data[baseHeader];
        if (!base) {
          return;
        }

        const time = data.Time;
        const timestamp = time ? new Date(time).getTime().toString() : "";
        const selectedColumns = {
          [baseHeader]: base,
          [TIMESTAMP_HEADER]: timestamp,
        };
        collection.push(selectedColumns);
      })
      .on("end", () => {
        resolve(collection);
      })
      .on("error", reject);
  });
};

const streamAppendMetrics = async (
  metric: IMetric,
  baseHeader: string,
  preRecords: any[]
): Promise<any[]> => {
  let count = 0;
  let index = 0;
  const records = [...preRecords];

  return new Promise(function (resolve, reject) {
    fs.createReadStream(metric.path)
      .pipe(csv({ separator: ",", headers: [baseHeader, ...metric.headers] }))
      .on("data", (data: any) => {
        if (count === 0) {
          count++;
          return;
        }

        const { [baseHeader]: base, ...rest } = data;
        if (!base) {
          return;
        }

        const selectedCollection = records[index];
        if (selectedCollection[baseHeader] === base) {
          records[index] = { ...selectedCollection, ...rest };
        }

        index++;
      })
      .on("end", () => {
        resolve(records);
      })
      .on("error", reject);
  });
};

export const streamMergeMetrics = async (
  metrics: IMetric[],
  newFile: string,
  baseHeader: string = TIME_HEADER
): Promise<void> => {
  if (metrics.length === 0) return;
  let collection = await streamGetBaseCollection(metrics[0], baseHeader);

  for (let i = 0; i < metrics.length; i++) {
    collection = await streamAppendMetrics(metrics[i], baseHeader, collection);

    if (i === metrics.length - 1) {
      var result = parse(collection);
      fs.writeFileSync(newFile, result);
    }
  }
};
