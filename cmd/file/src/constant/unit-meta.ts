import { MetricType } from "../enum/metrics-type";
import { UnitType } from "../enum/unit-type";
import { IUnit } from "../interfaces/unit";
import {
  CPU_HEADER,
  LATENCY_HEADER,
  MEM_HEADER,
  NET_HEADER,
  RESPONSE_HEADER,
  TPS_HEADER,
} from "./metrics-header";

export const UNIT_METADATA: IUnit[] = [
  {
    type: MetricType.CPU,
    unit: UnitType.Percentage,
    headers: CPU_HEADER,
  },
  {
    type: MetricType.Memory,
    unit: UnitType.Mebibytes,
    headers: MEM_HEADER,
  },
  {
    type: MetricType.Network,
    unit: UnitType.MegabytePerSecond,
    headers: NET_HEADER,
  },
  {
    type: MetricType.TPS,
    unit: UnitType.RequestPerSecond,
    headers: TPS_HEADER,
  },
  {
    type: MetricType.Latency,
    unit: UnitType.Millisecond,
    headers: LATENCY_HEADER,
  },
  {
    type: MetricType.Response,
    unit: UnitType.Millisecond,
    headers: RESPONSE_HEADER,
  },
];
