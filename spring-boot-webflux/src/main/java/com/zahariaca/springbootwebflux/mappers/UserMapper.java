package com.zahariaca.springbootwebflux.mappers;

import com.zahariaca.springbootwebflux.dto.UserDto;
import com.zahariaca.springbootwebflux.entity.UserEntity;
import com.zahariaca.springbootwebflux.model.UserModel;
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
