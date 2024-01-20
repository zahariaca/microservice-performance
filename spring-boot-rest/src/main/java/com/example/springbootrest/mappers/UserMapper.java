package com.example.springbootrest.mappers;

import com.example.springbootrest.dto.UserDto;
import com.example.springbootrest.entity.UserEntity;
import com.example.springbootrest.model.UserModel;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper(componentModel = "spring")
public interface UserMapper {

    UserModel dtoToModel(UserDto dto);

    @Mapping(target = "id", ignore = true)
    UserEntity modelToEntity(UserModel model);

    UserModel entityToModel(UserEntity entity);

    UserDto modelToDto(UserModel dto);

}
