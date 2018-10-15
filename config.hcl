Service Issues {
  Command Edit {
    Args = ["Owner", "Repo", "Number"]
  }
  Command Create {}
  Command Lock {}
  Command AddLabelsToIssue {
    RoutesName = "add-labels"
    Args = ["Owner", "Repo", "Number", "Labels"]
  }
}

Service Organizations {
  RouteName = "orgs"
  Command Get {}
}
