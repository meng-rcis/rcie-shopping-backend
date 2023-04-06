export const findNearestFailedTime = (
  timeStamp: string,
  failedResponse: any[]
): string => {
  let diff = 0;
  const timeStampInt = Number(timeStamp);

  failedResponse.forEach((error) => {
    if (timeStampInt > Number(error.timeStamp)) {
      return diff;
    }
    diff = Number(error.timeStamp) - timeStampInt;
  });

  return diff.toString();
};
