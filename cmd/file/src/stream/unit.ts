import fs from "fs";
import csv from "csv-parser";
// @ts-ignore
import { parse } from "json2csv";
import { IUnit } from "../interfaces/unit";
import { convertUnit } from "../utils/unit";
import {
  TIMESTAMP_HEADER,
  TIME_HEADER,
  TIME_TO_FAIL_HEADER,
} from "../constant/metrics-header";

export const streamStandardizeUnit = async (
  currentFile: string,
  newFile: string,
  unitMeta: IUnit[]
): Promise<void> => {
  let count = 0;
  const collection: any[] = [];
  const metricsHeaders = unitMeta.flatMap((unit) => unit.headers);

  return new Promise(function (resolve, reject) {
    fs.createReadStream(currentFile)
      .pipe(
        csv({
          separator: ",",
          headers: [
            TIME_HEADER,
            TIMESTAMP_HEADER,
            ...metricsHeaders,
            TIME_TO_FAIL_HEADER,
          ],
        })
      )
      .on("data", (data: any) => {
        if (count === 0) {
          count++;
          return;
        }

        const standardizedRow = convertUnit(data, unitMeta);
        collection.push(standardizedRow);
      })
      .on("end", () => {
        var result = parse(collection);
        fs.writeFileSync(newFile, result);
        resolve();
      })
      .on("error", reject);
  });
};
