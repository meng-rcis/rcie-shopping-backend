import { UnitType } from "../enum/unit-type";
import { IUnit } from "../interfaces/unit";

const findConvertRatio = (actualUnit: string, desiredUnit: string): number => {
  switch (actualUnit) {
    case UnitType.KilobytePerSecond:
      if (desiredUnit === UnitType.MegabytePerSecond) return 1 / 1000;
      if (desiredUnit === UnitType.GigabytePerSecond) return 1 / 1000000;
      return 1;
    case UnitType.MegabytePerSecond:
      if (desiredUnit === UnitType.KilobytePerSecond) return 1000;
      if (desiredUnit === UnitType.GigabytePerSecond) return 1 / 1000;
      return 1;
    case UnitType.GigabytePerSecond:
      if (desiredUnit === UnitType.KilobytePerSecond) return 1000000;
      if (desiredUnit === UnitType.MegabytePerSecond) return 1000;
    case UnitType.Second:
      if (desiredUnit === UnitType.Millisecond) return 1000;
    case UnitType.Millisecond:
      if (desiredUnit === UnitType.Second) return 1 / 1000;
    default:
      return -1;
  }
};

const convert = (value: string, desiredUnit: string): string => {
  let substrings = value.split(" ");
  const isSubstringIncluded = substrings.length > 0;
  if (!isSubstringIncluded) {
    console.log("Substrings not found: " + value);
    return value;
  }

  const isPercentageType = value.includes(UnitType.Percentage);
  if (isPercentageType && desiredUnit !== UnitType.Percentage) {
    console.log("Incorrect unit input, value: " + value);
    return value;
  }

  if (isPercentageType) {
    const number = parseFloat(value.replace("%", ""));
    const decimal = number / 100;
    return decimal.toString();
  }

  const number = parseFloat(substrings[0]);
  const actualUnit = substrings[1];
  if (actualUnit === desiredUnit) {
    return number.toString();
  }

  const ratio = findConvertRatio(actualUnit, desiredUnit);
  if (ratio === -1) {
    console.log("Cannot find unit of " + value);
    return value;
  }

  const convertedValue = number * ratio;
  return convertedValue.toString();
};

export const convertUnit = (data: any, unitMeta: IUnit[]): any => {
  const { Time, Timestamp, TimeToFail, ...rest } = data;
  const timestampHeader = "Timestamp [ms]";
  const standardizedRow: any = { Time, [timestampHeader]: Timestamp };

  unitMeta.forEach((meta) => {
    const { headers, unit } = meta;
    headers.forEach((header) => {
      const value = rest[header];

      if (value) {
        const convertedValue = convert(value, unit);
        const displayedUnit = unit === UnitType.Percentage ? "Ratios" : unit;
        const headerWithUnit = header + " [" + displayedUnit + "]";
        standardizedRow[headerWithUnit] = convertedValue;
      }
    });
  });

  const timeToFailHeader = "TimeToFail [ms]";
  standardizedRow[timeToFailHeader] = TimeToFail;
  return standardizedRow;
};
