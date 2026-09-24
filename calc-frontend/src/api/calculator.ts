const StatusCodes = {
  Success: 0,
  IncompleteInput: 1,
  DivisionByZero: 2,
  NegativeSqrt: 3,
} as const;

type CalculationStatusCode = (typeof StatusCodes)[keyof typeof StatusCodes];

interface Calculation {
  status_code: CalculationStatusCode;
  expression: string;
  result: string;
  error: string;
}

export async function calculate(expression: string): Promise<string | null> {
  const res = await fetch("/api/calculate", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ expression }),
  });

  if (res.status == 400) {
    return null;
  }

  if (!res.ok) {
    throw new Error(
      `Failed to calculate expressoin ${expression} with status ${res.status}`,
    );
  }

  return res.json().then((calculation: Calculation) => {
    if (calculation.status_code == StatusCodes.Success) {
      return calculation.result;
    }
    return calculation.error;
  });
}
