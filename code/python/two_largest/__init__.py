def two_largest(*nums):
  if nums[1] > nums[0]:
    if nums[2] > nums[1]:
      return [nums[2], nums[1]]
    else:
      return [nums[1], max(nums[2], nums[0])]
  else:
    if nums[0] > nums[2]:
      return [nums[0], max(nums[2], nums[1])]
    else:
      return [nums[2], nums[0]]
