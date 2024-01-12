#! This script preprocess myFitnessPal csv data given by user to something useful which grafana can display
import pandas as pd
import json

#!! Remeber to make a function to add column name to myFitnessPal csv data. Column should be ['Index', 'Date', 'Individual-Intake', 'Total-Intake'].
#!! Add a validation function for dataframe using pydantic


def AddCloumnHeaders(df: pd.DataFrame):
    return df.rename(
        columns={
            df.columns[0]: "Index",  # type: ignore
            df.columns[1]: "Date",  # type: ignore
            df.columns[2]: "Individual-Intake",  # type: ignore
            df.columns[3]: "Total-Intake",  # type: ignore
        }
    )


def ExtractTotalCalories(df: pd.DataFrame):
    new_df = pd.DataFrame(columns=["Date", "TotalCalories", "GoalCalories"])
    for _, val in df.iterrows():
        TotalCal, GoalCal = -1, -1
        temp_total = json.loads(val["Total-Intake"])  # type:ignore

        for i in temp_total["total"]:
            if i["name"] == "Calories":
                TotalCal = i["value"]

        for i in temp_total["goal"]:
            if i["name"] == "Calories":
                GoalCal = i["value"]

        new_entry = [val["Date"], TotalCal, GoalCal]

        new_df.loc[len(new_df.index)] = new_entry

    return new_df


def main():
    df = pd.read_csv("./data.csv")
    new_df = ExtractTotalCalories(df)
    new_df.to_csv("./output.csv")


if __name__ == "__main__":
    main()
